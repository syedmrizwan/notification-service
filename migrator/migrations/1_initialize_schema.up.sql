create table if not exists notification_handlers
(
	id serial not null
		constraint notification_handler_pk
			primary key,
	name text,
	rate_per_minute integer
);

alter table notification_handlers owner to postgres;

create table if not exists notification_texts
(
	id serial not null
		constraint notification_text_pk
			primary key,
	message text
);

alter table notification_texts owner to postgres;

create table if not exists notifications
(
	id serial not null
		constraint notification_pk
			primary key,
	priority text,
	user_id integer,
	notification_text_id integer
		constraint notification_notification_text__fk
			references notification_texts,
	notification_handler_id integer
		constraint notification_notification_handler__fk
			references notification_handlers,
	created_at timestamp default now()
);

alter table notifications owner to postgres;


insert into notification_handlers (id, name, rate_per_minute) values (1, 'SMS', 150);
insert into notification_handlers (id, name, rate_per_minute) values (2, 'Email', 300);