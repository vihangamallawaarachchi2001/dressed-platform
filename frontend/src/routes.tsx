// src/routes.tsx
import { Routes as RouterRoutes, Route, Navigate } from 'react-router-dom';
import { getRole, getToken } from './utils/auth';

// Public Pages
import Landing from './pages/Landing';
import Login from './pages/Login';
import Register from './pages/Register';

// Protected Pages
import DesignerDashboard from './pages/designer/Dashboard';
import SupplierDashboard from './pages/supplier/Dashboard';
import AdminDashboard from './pages/admin/Dashboard';
import type { JSX } from 'react';

// Role-based route wrappers
const ProtectedRoute = ({ children, allowedRoles }: { children: JSX.Element; allowedRoles: string[] }) => {
  const token = getToken();
  const role = getRole();

  if (!token) {
    return <Navigate to="/login" />;
  }

  if (!allowedRoles.includes(role || '')) {
    return <Navigate to="/" />;
  }

  return children;
};

const AuthRoute = ({ children }: { children: JSX.Element }) => {
  const token = getToken();
  return token ? <Navigate to={getRole() === 'designer' ? '/designer/dashboard' : '/supplier/dashboard'} /> : children;
};

const Routes = () => {
  return (
    <RouterRoutes>
      {/* Public Routes */}
      <Route path="/" element={<Landing />} />
      <Route path="/login" element={<AuthRoute><Login /></AuthRoute>} />
      <Route path="/register/:role" element={<AuthRoute><Register /></AuthRoute>} />

      {/* Designer Routes */}
      <Route
        path="/designer/dashboard"
        element={
          <ProtectedRoute allowedRoles={['designer']}>
            <DesignerDashboard />
          </ProtectedRoute>
        }
      />

      {/* Supplier Routes */}
      <Route
        path="/supplier/dashboard"
        element={
          <ProtectedRoute allowedRoles={['supplier']}>
            <SupplierDashboard />
          </ProtectedRoute>
        }
      />

      {/* Admin Routes */}
      <Route
        path="/admin"
        element={
          <ProtectedRoute allowedRoles={['admin']}>
            <AdminDashboard />
          </ProtectedRoute>
        }
      />

      {/* Fallback */}
      <Route path="*" element={<Navigate to="/" />} />
    </RouterRoutes>
  );
};

export default Routes;