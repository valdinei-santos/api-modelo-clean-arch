CREATE TABLE cliente (
    cpf          VARCHAR2(11),
    nm_cliente   VARCHAR2(70) NOT NULL,
    dt_nasc      DATE
);
alter table cliente add constraint cliente_pk primary key (cpf);

CREATE TABLE telefone (
    cpf       VARCHAR2(11) NOT NULL,
    numero    VARCHAR2(11) NOT NULL
);
alter table telefone add constraint telefone_pk  primary key (cpf, numero);
alter table telefone add constraint telefone_fk1 foreign key (cpf) references cliente (cpf);

-- Teste
insert into cliente values (1, 'Cliente 1', '02/07/1975');
insert into telefone values (1, '48999448383');
insert into telefone values (1, '4832215526');