import type { Design } from "../types";

// src/services/supplierService.ts
const API_BASE = '/api';

export interface SubmitQuoteRequest {
  design_id: string;
  price: number;
  eta_days: number;
  notes: string;
}

export const getPublicDesigns = async (): Promise<Design[]> => {
  const response = await fetch(`${API_BASE}/designs`);
  if (!response.ok) throw new Error('Failed to fetch designs');
  return response.json();
};

export const submitQuote = async (data: SubmitQuoteRequest): Promise<void> => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}/suppliers/quotes`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    const error = await response.json().catch(() => ({}));
    throw new Error(error.error || 'Failed to submit quote');
  }
};