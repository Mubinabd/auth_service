begin transaction;
    update books set name = 'Sariq devni minib' where id = 2;
    insert into books(name, page, student_id) values('Iymon', 456, 5);
commit;


begin transaction;
    update books set is_sold = true where id = 3;
commit;


begin transaction;
insert into 
transaction(id, amount, type, description, payment_type, from_card, to_card)
values('13a77bad-3bb4-4676-bb2e-dc987c5a4a87', 
4000, 'debit', 'desc', 'transfer', '5028ecc1-a730-4cfa-b4e0-4bec690504c9', null);
insert into 
transaction(id, amount, type, description, payment_type, from_card, to_card)
values('13a77bad-3bb4-4676-bb2e-dc987c5a4a87', 
4000, 'credit', 'desc 2', 'transfer', '5028ecc1-a730-4cfa-b4e0-4bec690504c9', null);
commit;