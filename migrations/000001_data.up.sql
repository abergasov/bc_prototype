CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    
create table circle_members
(
    circle_id uuid,
    member_id uuid,
    address   varchar,
    role      integer,
    constraint circle_members_pk
        primary key (circle_id, member_id)
);

INSERT INTO circle_members (circle_id, member_id, address, role) VALUES ('372e3c13-a552-4ebe-96fb-7f1bf45cf68c', 'e9c00793-0e0a-43d2-af68-c75e2303b86e', null, 3);
INSERT INTO circle_members (circle_id, member_id, address, role) VALUES ('372e3c13-a552-4ebe-96fb-7f1bf45cf68c', 'b2e88baa-6b6d-4d17-8e4d-4a645f223c0c', null, 1);
INSERT INTO circle_members (circle_id, member_id, address, role) VALUES ('372e3c13-a552-4ebe-96fb-7f1bf45cf68c', 'edb94aad-8266-4d06-a08a-644bf100b229', null, 1);

create table circle_contract
(
    circle_id        uuid
        constraint circle_contract_pk
            primary key,
    contract_address varchar
);

alter table circle_members
    add constraint circle_members_pk2
        unique (address);