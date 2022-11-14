create table services(
    id bigserial primary key,
    name varchar(250),
    description text,
    active boolean not null default true,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);
create table countries(
                          id bigserial primary key,
                          name text,
                          active boolean not null default true,
                          created_at timestamptz default current_timestamp,
                          updated_at timestamptz not null default current_timestamp,
                          deleted_at timestamptz not null default current_timestamp
);
create table cities(
    id bigserial primary key,
    name varchar(250),
    country_id int references countries,
    active boolean not null default true,
    created_at timestamptz default current_timestamp,
    deleted_at timestamptz not null default current_timestamp
);

create table feedbacks(
    id bigserial primary key,
    user_id bigint not null,
    city_id int references cities,
    massage text not null,
    photo text,
    created_at timestamptz default current_timestamp
);