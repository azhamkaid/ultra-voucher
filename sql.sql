CREATE TABLE student (id int ,name varchar,parent_id int);

INSERT INTO student (id,name,parent_id)VALUES(1,'zaki',2),(2,'ilham',null),(3,'irwan',2),(4,'arka',3);

SELECT * from student;
SELECT 
    c.id AS id,
    c.name AS name,
    p.name AS parent_name
FROM 
    student AS c
left JOIN 
    student AS p ON c.parent_id = p.id
order by id;