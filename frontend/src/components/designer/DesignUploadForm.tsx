/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { uploadDesign } from '../../services/designService';

const DesignUploadForm = () => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [category, setCategory] = useState('Women');
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!file) {
      setError('Please select a file');
      return;
    }

    const formData = new FormData();
    formData.append('title', title);
    formData.append('description', description);
    formData.append('category', category);
    formData.append('file', file);

    setUploading(true);
    setError('');
    try {
      await uploadDesign(formData);
      alert('Design uploaded! Go to "My Designs" to submit it.');
    } catch (err: any) {
      setError(err.message);
    } finally {
      setUploading(false);
    }
  };

  const categories = ['Men', 'Women', 'Boy', 'Girl', 'Unisex'];

  return (
    <div className="bg-white p-6 rounded-xl shadow-md">
      <h2 className="text-xl font-bold mb-4">Upload New Design</h2>
      {error && <div className="bg-red-100 text-red-700 p-2 rounded mb-4">{error}</div>}
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-gray-700 mb-1">Title</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded"
            required
          />
        </div>
        <div>
          <label className="block text-gray-700 mb-1">Description</label>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded"
            rows={3}
            required
          />
        </div>
        <div>
          <label className="block text-gray-700 mb-1">Category</label>
          <select
            value={category}
            onChange={(e) => setCategory(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded"
          >
            {categories.map((cat) => (
              <option key={cat} value={cat}>{cat}</option>
            ))}
          </select>
        </div>
        <div>
          <label className="block text-gray-700 mb-1">Design File (Image/PDF)</label>
          <input
            type="file"
            accept="image/*,.pdf"
            onChange={(e) => setFile(e.target.files?.[0] || null)}
            className="w-full"
            required
          />
        </div>
        <button
          type="submit"
          disabled={uploading}
          className="w-full bg-indigo-600 text-white py-2 rounded hover:bg-indigo-700 disabled:opacity-50"
        >
          {uploading ? 'Uploading...' : 'Upload Design'}
        </button>
      </form>
    </div>
  );
};

export default DesignUploadForm;