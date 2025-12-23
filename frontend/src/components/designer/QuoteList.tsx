/* eslint-disable @typescript-eslint/no-explicit-any */
// src/components/designer/QuoteList.tsx
import { useState, useEffect } from 'react';
import { getMyQuotes } from '../../services/quoteService';
import QuoteItem from './QuoteItem';
import type { Quote } from '../../types';

const QuoteList = () => {
  const [quotes, setQuotes] = useState<Quote[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchQuotes = async () => {
      try {
        const data = await getMyQuotes();
        setQuotes(data);
      } catch (err: any) {
        alert('Failed to load quotes: ' + err.message);
      } finally {
        setLoading(false);
      }
    };
    fetchQuotes();
  }, []);

  return (
    <div className="bg-white p-6 rounded-xl shadow-md">
      <h2 className="text-xl font-bold mb-4">Incoming Quotes</h2>
      {loading ? (
        <p>Loading quotes...</p>
      ) : quotes.length === 0 ? (
        <p className="text-gray-500">No quotes yet. Submit designs to receive quotes!</p>
      ) : (
        <div>
          {quotes.map((quote) => (
            <QuoteItem key={quote.id} quote={quote} />
          ))}
        </div>
      )}
    </div>
  );
};

export default QuoteList;