create table if not exists notification_handler
(
	id serial not null
		constraint notification_handler_pk
			primary key,
	name text,
	rate_per_minute integer
);

alter table notification_handler owner to postgres;

create table if not exists notification_text
(
	id serial not null
		constraint notification_text_pk
			primary key,
	message text
);

alter table notification_text owner to postgres;

create table if not exists notification
(
	id serial not null
		constraint notification_pk
			primary key,
	priority text,
	status text,
	user_id integer,
	notification_text_id integer
		constraint notification_notification_text__fk
			references notification_text,
	notification_handler_id integer
		constraint notification_notification_handler__fk
			references notification_handler,
	created_at timestamp default now()
);

alter table notification owner to postgres;


insert into notification_handler (id, name, rate_per_minute) values (1, 'SMS', 15);
insert into notification_handler (id, name, rate_per_minute) values (2, 'Email', 30);