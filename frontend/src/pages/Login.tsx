/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { login } from '../services/authService';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    try {
      const { access_token, role } = await login({ email, password });
      localStorage.setItem('token', access_token);
      localStorage.setItem('role', role);

      if (role === 'designer') {
        navigate('/designer/dashboard');
      } else if (role === 'supplier') {
        navigate('/supplier/dashboard');
      } else if (role === 'admin') {
        navigate('/admin');
      }
    } catch (err: any) {
      setError(err.message || 'Invalid email or password');
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center p-4">
      <div className="bg-white p-8 rounded-xl shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold text-center mb-6">Sign In to Dressed™</h2>
        {error && <div className="bg-red-100 text-red-700 p-3 rounded mb-4">{error}</div>}

        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700 mb-2">Email</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700 mb-2">Password</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-indigo-600 text-white py-2 rounded-lg hover:bg-indigo-700 transition"
          >
            Sign In
          </button>
        </form>

        <div className="mt-6 text-center text-gray-600">
          <p>
            Don’t have an account?{' '}
          </p>
          <div className='flex gap-4 my-5 items-center justify-center w-full'>
            <div className='mx-2'>
              <Link to="/register/designer" className="text-white hover:underline bg-indigo-600 rounded-[5px] px-2 py-2 ">
              Sign up as Designer
            </Link>
            </div>
           <div>
             <Link to="/register/supplier" className="text-indigo-600 border border-indigo-600 rounded-[5px] px-2 py-2 hover:underline">
              Sign up as Supplier
            </Link>
           </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;