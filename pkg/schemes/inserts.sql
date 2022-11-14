insert into countries(name)
values ('Tajikistan'),
       ('Russian');

insert into cities(name)
values ('Dushanbe'),
       ('Khujand');


insert into country_cities(country_id, city_id)
values (1, 1),
       (1,2);


insert into services(name, description)
values ('Books', 'More sale'),
       ('Foods', 'pizza');


insert into local_services(country_id, city_id, service_id)
values(1, 1, 1),
      (1,1,2)