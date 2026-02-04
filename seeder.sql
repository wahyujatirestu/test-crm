INSERT INTO membership (
  name, password, address, is_active, created_by
) VALUES (
  'Admin',
  md5('password123'),
  'Jakarta',
  true,
  'seeder'
);

INSERT INTO contact (
  membership_id, contact_type, contact_value, is_active, created_by
) VALUES (
  1,
  'email',
  'admin@mail.com',
  true,
  'seeder'
);
