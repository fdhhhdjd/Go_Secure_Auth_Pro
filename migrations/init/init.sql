-- Create the database
CREATE DATABASE ${POSTGRES_DB};

-- Switch to the new database
\c ${POSTGRES_DB};

-- Create the tables
\i migrations/1_create_table_user.sql
\i migrations/2_create_table_password_history.sql
\i migrations/3_create_table_devices.sql
\i migrations/4_create_table_social_logins.sql
\i migrations/5_create_table_otp.sql