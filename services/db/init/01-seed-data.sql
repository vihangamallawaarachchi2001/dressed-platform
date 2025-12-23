-- Seed realistic sample data for Dressed™ platform
-- Safe to run multiple times (idempotent)

DO $$
DECLARE
  designer1_id TEXT := 'designer-001';
  designer2_id TEXT := 'designer-002';
  supplier1_id TEXT := 'supplier-001';
  supplier2_id TEXT := 'supplier-002';
BEGIN
  -- === USERS (Designers & Suppliers) ===
  INSERT INTO users (id, email, password, role, created_at)
  VALUES
    -- Designer 1
    (designer1_id, 'designer1@dressed.com', '$2a$12$0cJ7KzX6y7qQvT9dJv8Z.eXqF1dJv8Z.eXqF1dJv8Z.eXqF1d', 'designer', NOW()),
    -- Designer 2
    (designer2_id, 'designer2@dressed.com', '$2a$12$0cJ7KzX6y7qQvT9dJv8Z.eXqF1dJv8Z.eXqF1dJv8Z.eXqF1d', 'designer', NOW()),
    -- Supplier 1
    (supplier1_id, 'supplier1@dressed.com', '$2a$12$0cJ7KzX6y7qQvT9dJv8Z.eXqF1dJv8Z.eXqF1dJv8Z.eXqF1d', 'supplier', NOW()),
    -- Supplier 2
    (supplier2_id, 'supplier2@dressed.com', '$2a$12$0cJ7KzX6y7qQvT9dJv8Z.eXqF1dJv8Z.eXqF1dJv8Z.eXqF1d', 'supplier', NOW())
  ON CONFLICT (email) DO NOTHING;

  -- === SUPPLIER PROFILES ===
  INSERT INTO suppliers (id, user_id, company_name, description, capabilities, status, availability, created_at)
  VALUES
    (
      'supplier-profile-001',
      supplier1_id,
      'EcoTextiles Co.',
      'Sustainable fabric specialist with GOTS certification',
      'Organic Cotton, Hemp, Recycled Polyester',
      'ACTIVE',
      'AVAILABLE',
      NOW()
    ),
    (
      'supplier-profile-002',
      supplier2_id,
      'Urban Stitch Labs',
      'Fast-turnaround digital printing and small-batch production',
      'Digital Print, Cut & Sew, Sample Making',
      'ACTIVE',
      'BUSY',
      NOW()
    )
  ON CONFLICT (user_id) DO NOTHING;

  -- === DESIGNS (6 total) ===
  INSERT INTO designs (id, designer_id, title, description, category, file_path, status, created_at, updated_at)
  VALUES
    -- Designer 1: Women's dress (SUBMITTED)
    ('design-001', designer1_id, 'Eco Linen Maxi Dress', 'Breathable summer dress made from 100% linen', 'Women', '/uploads/linen_dress.jpg', 'SUBMITTED', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),

    -- Designer 1: Men's shirt (SUBMITTED)
    ('design-002', designer1_id, 'Recycled Cotton Oxford Shirt', 'Classic fit shirt from post-consumer cotton', 'Men', '/uploads/oxford_shirt.jpg', 'SUBMITTED', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),

    -- Designer 2: Kids unisex set (SUBMITTED)
    ('design-003', designer2_id, 'Organic Cotton Kids Set', 'Gender-neutral t-shirt and pants for toddlers', 'Unisex', '/uploads/kids_set.jpg', 'SUBMITTED', NOW() - INTERVAL '3 hours', NOW() - INTERVAL '3 hours'),

    -- Designer 2: Boy's jacket (DRAFT - not visible to suppliers)
    ('design-004', designer2_id, 'Waterproof Boys Rain Jacket', 'Lightweight jacket with recycled lining', 'Boy', '/uploads/rain_jacket.jpg', 'DRAFT', NOW(), NOW()),

    -- Designer 1: Girl's dress (SUBMITTED)
    ('design-005', designer1_id, 'Floral Print Party Dress', 'Cotton blend with ruffled sleeves', 'Girl', '/uploads/party_dress.jpg', 'SUBMITTED', NOW() - INTERVAL '5 hours', NOW() - INTERVAL '5 hours'),

    -- Designer 2: Women's top (DRAFT)
    ('design-006', designer2_id, 'Linen Crop Top', 'Summer essential with adjustable straps', 'Women', '/uploads/crop_top.jpg', 'DRAFT', NOW(), NOW())
  ON CONFLICT (id) DO NOTHING;

  -- === QUOTES (4 total) ===
  INSERT INTO quotes (id, design_id, designer_id, supplier_id, price, eta_days, notes, status, created_at)
  VALUES
    -- Quote for design-001 (Women's dress) from Supplier 1 → ACCEPTED
    (
      'quote-001',
      'design-001',
      designer1_id,
      supplier1_id,
      89.99,
      14,
      'Can produce in organic linen as specified. MOQ: 50 units.',
      'ACCEPTED',
      NOW() - INTERVAL '1 day'
    ),

    -- Quote for design-001 from Supplier 2 → QUOTED (pending)
    (
      'quote-002',
      'design-001',
      designer1_id,
      supplier2_id,
      75.50,
      10,
      'Faster turnaround using digital print on linen blend. Price includes sample.',
      'QUOTED',
      NOW() - INTERVAL '12 hours'
    ),

    -- Quote for design-002 (Men's shirt) from Supplier 1 → QUOTED
    (
      'quote-003',
      'design-002',
      designer1_id,
      supplier1_id,
      45.00,
      21,
      'Minimum order 100 units. Lead time includes GOTS certification.',
      'QUOTED',
      NOW() - INTERVAL '6 hours'
    ),

    -- Quote for design-003 (Kids set) from Supplier 2 → ACCEPTED
    (
      'quote-004',
      'design-003',
      designer2_id,
      supplier2_id,
      22.75,
      7,
      'Perfect for our small-batch production line. Can deliver in one week.',
      'ACCEPTED',
      NOW() - INTERVAL '2 hours'
    )
  ON CONFLICT (id) DO NOTHING;

  -- === ORDERS (2 total, from accepted quotes) ===
  INSERT INTO orders (id, quote_id, designer_id, supplier_id, status, created_at)
  VALUES
    ('order-001', 'quote-001', designer1_id, supplier1_id, 'CREATED', NOW() - INTERVAL '1 day'),
    ('order-002', 'quote-004', designer2_id, supplier2_id, 'CREATED', NOW() - INTERVAL '2 hours')
  ON CONFLICT (id) DO NOTHING;

  -- === PAYMENTS (2 total, linked to orders) ===
  INSERT INTO payments (id, order_id, amount, status, created_at)
  VALUES
    ('payment-001', 'order-001', 89.99, 'SUCCESS', NOW() - INTERVAL '1 day'),
    ('payment-002', 'order-002', 22.75, 'PENDING', NOW() - INTERVAL '2 hours')
  ON CONFLICT (id) DO NOTHING;

  RAISE NOTICE '✅ Dressed™ sample data seeded successfully';
END $$;