// src/pages/Landing.tsx
import { Link } from 'react-router-dom';

const Landing = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
      {/* Header */}
      <header className="px-6 py-4">
        <div className="container mx-auto flex justify-between items-center">
          <h1 className="text-2xl font-bold text-indigo-700">Dressed™</h1>
          <nav>
            <Link
              to="/login"
              className="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition"
            >
              Sign In
            </Link>
          </nav>
        </div>
      </header>

      {/* Hero */}
      <section className="py-16 bg-white">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Where Fashion Design Meets Manufacturing
          </h1>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto mb-10">
            Dressed™ bridges the gap between creative designers and reliable suppliers —
            streamlining the journey from concept to production.
          </p>
          <div className="flex justify-center gap-4">
            <Link
              to="/register/designer"
              className="px-6 py-3 bg-indigo-600 text-white rounded-lg font-medium hover:bg-indigo-700 transition"
            >
              Join as Designer
            </Link>
            <Link
              to="/register/supplier"
              className="px-6 py-3 bg-white text-indigo-600 border border-indigo-600 rounded-lg font-medium hover:bg-indigo-50 transition"
            >
              Join as Supplier
            </Link>
          </div>
        </div>
      </section>

      {/* How It Works */}
      <section className="py-16 bg-gray-50">
        <div className="container mx-auto px-6">
          <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">How It Works</h2>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="bg-white p-6 rounded-lg shadow-sm">
              <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-4">
                <span className="text-blue-600 font-bold">1</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">Designers Upload</h3>
              <p className="text-gray-600">
                Submit your clothing designs with images/PDFs, categorized by gender or style.
              </p>
            </div>
            <div className="bg-white p-6 rounded-lg shadow-sm">
              <div className="w-12 h-12 bg-emerald-100 rounded-full flex items-center justify-center mb-4">
                <span className="text-emerald-600 font-bold">2</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">Suppliers Quote</h3>
              <p className="text-gray-600">
                Review designs and submit quotes with pricing, delivery time, and notes.
              </p>
            </div>
            <div className="bg-white p-6 rounded-lg shadow-sm">
              <div className="w-12 h-12 bg-amber-100 rounded-full flex items-center justify-center mb-4">
                <span className="text-amber-600 font-bold">3</span>
              </div>
              <h3 className="text-xl font-semibold mb-2">Orders Fulfilled</h3>
              <p className="text-gray-600">
                Designers accept quotes, and orders move to production with secure payment.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Role Cards (as before, but styled consistently) */}
      <section className="py-16">
        <div className="container mx-auto px-6">
          <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">
            Choose Your Role
          </h2>
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 max-w-4xl mx-auto">
            {/* Designer */}
            <div className="bg-white p-8 rounded-xl shadow-md border border-gray-200">
              <h3 className="text-2xl font-bold text-blue-700 mb-4">For Designers</h3>
              <ul className="text-gray-600 mb-6 space-y-2">
                <li>✅ Upload designs (images or PDFs)</li>
                <li>✅ Categorize by Men, Women, Boy, Girl, Unisex</li>
                <li>✅ Receive and compare supplier quotes</li>
                <li>✅ Accept quotes to create orders</li>
              </ul>
              <Link
                to="/register/designer"
                className="inline-block px-6 py-3 bg-blue-600 text-white rounded-lg font-medium hover:bg-blue-700 transition"
              >
                Start Designing
              </Link>
            </div>

            {/* Supplier */}
            <div className="bg-white p-8 rounded-xl shadow-md border border-gray-200">
              <h3 className="text-2xl font-bold text-emerald-700 mb-4">For Suppliers</h3>
              <ul className="text-gray-600 mb-6 space-y-2">
                <li>✅ Browse all submitted designs</li>
                <li>✅ Submit quotes with price & ETA</li>
                <li>✅ Add notes explaining your offer</li>
                <li>✅ Track quote status (pending/accepted)</li>
              </ul>
              <Link
                to="/register/supplier"
                className="inline-block px-6 py-3 bg-emerald-600 text-white rounded-lg font-medium hover:bg-emerald-700 transition"
              >
                Start Quoting
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* Trust */}
      <section className="py-12 bg-indigo-900 text-white">
        <div className="container mx-auto px-6 text-center">
          <h2 className="text-2xl font-bold mb-4">Trusted by Creators & Manufacturers</h2>
          <p className="max-w-2xl mx-auto opacity-90">
            Join a growing community of fashion innovators and production experts.
          </p>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-800 text-white py-8">
        <div className="container mx-auto px-6 text-center">
          <p>© 2025 Dressed™. All rights reserved.</p>
          <p className="mt-2 text-gray-400 text-sm">
            Simplifying fashion production, one design at a time.
          </p>
        </div>
      </footer>
    </div>
  );
};

export default Landing;