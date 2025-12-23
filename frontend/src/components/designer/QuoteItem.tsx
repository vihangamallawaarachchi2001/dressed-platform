/* eslint-disable @typescript-eslint/no-explicit-any */
import type { Quote } from '../../types';
import { acceptQuote } from '../../services/quoteService';
import { useState } from 'react';

const QuoteItem = ({ quote }: { quote: Quote }) => {
  const [accepting, setAccepting] = useState(false);
  const [accepted, setAccepted] = useState(false);

  const handleAccept = async () => {
    if (accepting || accepted) return;
    setAccepting(true);
    try {
      await acceptQuote(quote.id);
      setAccepted(true);
      alert('Order created successfully!');
    } catch (err: any) {
      alert('Failed to accept quote: ' + err.message);
    } finally {
      setAccepting(false);
    }
  };

  return (
    <div className="border border-gray-200 p-4 rounded-lg mb-3">
      <div className="flex justify-between">
        <div>
          <p className="font-medium">Price: ${quote.price.toFixed(2)}</p>
          <p className="text-sm text-gray-600">ETA: {quote.eta_days} days</p>
        </div>
        <button
          onClick={handleAccept}
          disabled={accepted || accepting}
          className={`px-3 py-1 rounded text-sm ${
            accepted
              ? 'bg-green-100 text-green-800'
              : accepting
              ? 'bg-gray-100 text-gray-500'
              : 'bg-emerald-600 text-white hover:bg-emerald-700'
          }`}
        >
          {accepted ? 'Accepted' : accepting ? 'Accepting...' : 'Accept'}
        </button>
      </div>
      {quote.notes && (
        <div className="mt-2 p-3 bg-gray-50 rounded text-sm">
          <strong>Notes:</strong> {quote.notes}
        </div>
      )}
    </div>
  );
};

export default QuoteItem;