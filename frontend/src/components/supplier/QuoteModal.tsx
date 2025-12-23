/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { submitQuote } from '../../services/supplierService';

interface QuoteModalProps {
  designId: string;
  designTitle: string;
  onClose: () => void;
}

const QuoteModal = ({ designId, designTitle, onClose }: QuoteModalProps) => {
  const [price, setPrice] = useState('');
  const [eta, setEta] = useState('');
  const [notes, setNotes] = useState('');
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    const priceNum = parseFloat(price);
    const etaNum = parseInt(eta, 10);

    if (isNaN(priceNum) || priceNum <= 0) {
      setError('Price must be a positive number');
      return;
    }
    if (isNaN(etaNum) || etaNum <= 0) {
      setError('ETA must be a positive number of days');
      return;
    }

    setSubmitting(true);
    try {
      await submitQuote({
        design_id: designId,
        price: priceNum,
        eta_days: etaNum,
        notes,
      });
      alert('Quote submitted successfully!');
      onClose();
    } catch (err: any) {
      setError(err.message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-xl w-full max-w-md p-6">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-lg font-bold">Quote for: {designTitle}</h3>
          <button onClick={onClose} className="text-gray-500 hover:text-gray-700">
            ✕
          </button>
        </div>

        {error && <div className="bg-red-100 text-red-700 p-2 rounded mb-4">{error}</div>}

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-gray-700 text-sm mb-1">Price ($)</label>
            <input
              type="number"
              step="0.01"
              value={price}
              onChange={(e) => setPrice(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded"
              placeholder="99.99"
              required
            />
          </div>
          <div>
            <label className="block text-gray-700 text-sm mb-1">ETA (Days)</label>
            <input
              type="number"
              value={eta}
              onChange={(e) => setEta(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded"
              placeholder="7"
              required
            />
          </div>
          <div>
            <label className="block text-gray-700 text-sm mb-1">Notes (Optional)</label>
            <textarea
              value={notes}
              onChange={(e) => setNotes(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded"
              rows={3}
              placeholder="e.g., Can produce in organic cotton for +$5"
            />
          </div>
          <div className="flex gap-3">
            <button
              type="button"
              onClick={onClose}
              className="flex-1 px-4 py-2 border border-gray-300 rounded hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={submitting}
              className="flex-1 bg-emerald-600 text-white px-4 py-2 rounded hover:bg-emerald-700 disabled:opacity-50"
            >
              {submitting ? 'Submitting...' : 'Submit Quote'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default QuoteModal;