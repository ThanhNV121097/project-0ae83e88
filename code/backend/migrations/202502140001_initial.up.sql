create table if not exists schema_migrations (
  filename text primary key,
  applied_at timestamptz not null default now()
);

create table if not exists greetings (
  id smallint primary key,
  text text not null check (length(trim(text)) > 0),
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  constraint greetings_singleton check (id = 1)
);

insert into greetings (id, text)
values (1, 'Hello Word')
on conflict (id) do nothing;
