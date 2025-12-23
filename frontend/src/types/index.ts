export interface Design {
  id: string;
  title: string;
  description: string;
  category: string;
  filePath: string;
  status: string;
  createdAt: number;
  updatedAt: number;
}

export interface Quote {
  id: string;
  designId: string;
  designerId: string;
  supplierId: string;
  price: number;
  eta_days: number;
  notes: string;
  status: string;
  createdAt: string;
}