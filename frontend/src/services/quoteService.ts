import type { Quote } from "../types";

// src/services/quoteService.ts
const API_BASE = 'http://localhost:8000';

export const getMyQuotes = async (): Promise<Quote[]> => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}/quotes/my-quotes`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!response.ok) throw new Error('Failed to fetch quotes');
  return response.json();
};

export const acceptQuote = async (quoteId: string): Promise<void> => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}/quotes/${quoteId}/accept`, {
    method: 'POST',
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!response.ok) throw new Error('Failed to accept quote');
};