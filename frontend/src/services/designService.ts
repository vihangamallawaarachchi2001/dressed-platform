import type { Design } from "../types";

// src/services/designService.ts
const API_BASE = 'http://localhost:8000/designs';

export const uploadDesign = async (formData: FormData): Promise<void> => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}`, {
    method: 'POST',
    headers: { Authorization: `Bearer ${token}` },
    body: formData,
  });
  if (!response.ok) throw new Error('Failed to upload design');
};

export const getMyDesigns = async (): Promise<Design[]> => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}`, {
    method: 'GET',
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!response.ok) throw new Error('Failed to fetch designs');
  return response.json();
};

export const submitDesign = async (id: string): Promise<void> => {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}/${id}/submit`, {
    method: 'PATCH',
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!response.ok) throw new Error('Failed to submit design');
};