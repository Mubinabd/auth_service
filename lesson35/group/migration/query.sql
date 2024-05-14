CREATE TABLE student(
                        student_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        name VARCHAR(32),
                        lastname VARCHAR(32),
                        phone VARCHAR(13),
                        age INT,
                        created_at TIMESTAMP DEFAULT NOW(),
                        update_at TIMESTAMP ,
                        delated_at BIGINT DEFAULT 0
);

CREATE TABLE course(
                       course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       course_name VARCHAR(64),
                       created_at TIMESTAMP DEFAULT NOW(),
                       update_at TIMESTAMP ,
                       delated_at BIGINT DEFAULT 0
);

CREATE TABLE student_course(
                               student_course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               student_id UUID REFERENCES student(student_id),
                               course_id UUID REFERENCES course(course_id),
                               created_at TIMESTAMP DEFAULT NOW(),
                               update_at TIMESTAMP ,
                               delated_at BIGINT DEFAULT 0
);
CREATE UNIQUE INDEX idx_student_course_unique ON student_course (student_id, course_id);

CREATE TABLE grade (
                       grade_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       student_course_id UUID UNIQUE REFERENCES student_course(student_course_id),
                       grade FLOAT CHECK (grade >= 1 AND grade <= 5),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);



INSERT INTO student (student_id, name, lastname, phone, age)
VALUES
    ('94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', 'Aziz', 'Yusupov', '+998901234567', 20),
    ('c6928344-ddd1-49b8-b70d-dfaa89c129b5', 'Dilnoza', 'Kamilova', '+998901234568', 21),
    ('581ad352-0d11-4c91-a4c0-6c280d174b66', 'Farrukh', 'Zokirov', '+998901234569', 22),
    ('2253f666-9640-4291-a62a-e34479d11ec9', 'Gulnara', 'Rakhimova', '+998901234570', 23),
    ('9ce54ed8-f321-4f48-a492-63e41d287e4d', 'Ibrohim', 'Abdullayev', '+998901234571', 24),
    ('a67b3065-407c-46b0-b7e8-4e9ba4f71527', 'Jasur', 'Karimov', '+998901234572', 19),
    ('d9a7c912-7772-4ecd-aecc-e461a6e0fb97', 'Kamila', 'Safarova', '+998901234573', 20),
    ('cfc80055-02e3-4e1c-b510-ccf5832a008f', 'Lola', 'Tursunova', '+998901234574', 21),
    ('9e5916d9-7f64-449d-9cab-6fcbc5cc8b9a', 'Murod', 'Alimov', '+998901234575', 22),
    ('fd3e0268-7d0c-44f5-b508-8952e7cc6748', 'Nigora', 'Begimova', '+998901234576', 23),
    ('6fd9b5d1-4104-4b64-8240-9a473cf02d95', 'Otabek', 'Nazarov', '+998901234577', 24),
    ('b7d9d93a-7299-4da2-b288-e7c38740df17', 'Pulat', 'Hamidov', '+998901234578', 19),
    ('61b441bd-9465-41c3-a32e-38912c76fad6', 'Rustam', 'Sobirov', '+998901234579', 20),
    ('0666c4d6-cdb6-45c2-9aa7-53a3906b5dd4', 'Saida', 'Mirzaeva', '+998901234580', 21),
    ('52381e70-392a-4c0e-80c2-1cca456b4a5b', 'Tahir', 'Yuldoshev', '+998901234581', 22),
    ('7aaf2254-6e32-4b22-a06d-fa4a01acb92c', 'Ulugbek', 'Rustamov', '+998901234582', 23),
    ('f9030edc-de0a-4798-977f-19c05c61c75c', 'Vohid', 'Abdukarimov', '+998901234583', 24),
    ('851fb8ab-f298-4a17-bb93-5620853771a3', 'Yulduz', 'Usmanova', '+998901234584', 19),
    ('77a868b0-3d31-41e1-9885-5b39b7a7737c', 'Zafar', 'Ganiev', '+998901234585', 20),
    ('9335e840-27ae-4d60-afa0-b1c222cb047a', 'Bahrom', 'Yakubov', '+998901234586', 21),
    ('2dae324a-0272-4aba-ab62-fddc6077128d', 'Davron', 'Fayziev', '+998901234587', 22),
    ('30f2a30c-abc9-474a-af30-819f76aa3d95', 'Elyor', 'Sultonov', '+998901234588', 23),
    ('9b027e4e-22ce-4724-85fb-b3561f058d33', 'Feruza', 'Karimova', '+998901234589', 24),
    ('2081dba4-1010-40b8-8d82-b0022faa8d74', 'Giyos', 'Bekchanov', '+998901234590', 19),
    ('8aa2e3a7-e5e9-4900-8579-726edf93bd0c', 'Husan', 'Aliyev', '+998901234591', 20),
    ('93755307-5d53-4b62-b4eb-024cc9cfd6c7', 'Ismoil', 'Jurayev', '+998901234592', 21),
    ('aa16f2ad-4bee-423b-90dc-55f17380ce47', 'Jamshid', 'Rasulov', '+998901234593', 22),
    ('8c62875b-3d47-471d-ab2b-542d9627cfe0', 'Komil', 'Yusupov', '+998901234594', 23),
    ('9cd6f1ca-6608-4fa3-a41a-9a0bfb97d8f1', 'Laziz', 'Tursunov', '+998901234595', 24),
    ('06cef7a1-2afc-40d1-b149-13909243dd16', 'Malika', 'Hamidova', '+998901234596', 19);


INSERT INTO course (course_id, course_name)
VALUES
    ('eb4100eb-b697-43ef-9775-887ab192b0d3', 'Mathematics'),
    ('616655f0-d066-4ce1-94d5-7d49d95f0bd3', 'Physics'),
    ('42dfec60-e6df-4cd6-a1e6-2f228f827b5f', 'Chemistry'),
    ('7b534913-aeaa-4837-b617-f3e3c504f254', 'Biology'),
    ('e1e5be66-32ff-4ea7-95e4-3717d725bd7f', 'Computer Science');


INSERT INTO student_course (student_course_id, student_id, course_id)
VALUES
    ('94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', '94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),
    ('4b8d1eab-845b-4a4f-a6d7-bb30b4a9f8ef', '94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('7d29c143-cc48-4c12-b40f-51b3e9f7c04b', '94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('e7f87936-1d6b-4bfc-a7d0-686986ef3e6a', '94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('f02cb0c2-2758-44d2-b9e7-1d7863d32dcf', '94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),

    ('29f48565-d833-4753-b0a3-ced061fe717f', 'c6928344-ddd1-49b8-b70d-dfaa89c129b5', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),
    ('afeaa3e3-b9ff-48c4-b324-01e470c7d514', 'c6928344-ddd1-49b8-b70d-dfaa89c129b5', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('c284c145-75ff-4f6a-b7d9-83d1b9ed601d', 'c6928344-ddd1-49b8-b70d-dfaa89c129b5', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('4c3bb30d-b4cb-44a1-af46-37c48e46b50e', 'c6928344-ddd1-49b8-b70d-dfaa89c129b5', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),
    ('ac912f08-b3c5-4d9a-a3f8-d51a2a1e9b2b', 'c6928344-ddd1-49b8-b70d-dfaa89c129b5', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),

    ('8c5b3e05-df9d-4d3f-ae1f-2c476ee0d46b', '581ad352-0d11-4c91-a4c0-6c280d174b66', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('f3f5b9f2-4f1e-4d39-a1a2-16ac8975f3fb', '581ad352-0d11-4c91-a4c0-6c280d174b66', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('f87c0d5a-f876-4a79-94c6-ba29a1f3bf6f', '581ad352-0d11-4c91-a4c0-6c280d174b66', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('394cd112-05e0-44cb-a065-cc38b6cb67a2', '581ad352-0d11-4c91-a4c0-6c280d174b66', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),
    ('e1836f70-36b7-4854-b2d6-dce4e470e7cb', '581ad352-0d11-4c91-a4c0-6c280d174b66', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),

    ('4a32b0c5-8d33-4d18-a74a-b46c607c0e39', '2253f666-9640-4291-a62a-e34479d11ec9', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),
    ('15b2cb7a-5022-4b82-8a61-55bb0a3d0b9b', '2253f666-9640-4291-a62a-e34479d11ec9', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('f2543035-d27e-47ff-8307-3a51b1d722d8', '2253f666-9640-4291-a62a-e34479d11ec9', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('062f7a7b-8810-4899-b7b6-126a8f5360f2', '2253f666-9640-4291-a62a-e34479d11ec9', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('5534f6a2-d8ec-40d1-8045-8818a4e4d1d8', '2253f666-9640-4291-a62a-e34479d11ec9', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),

    ('7ee366eb-13bb-4fc2-a166-af218e90e37a', '9ce54ed8-f321-4f48-a492-63e41d287e4d', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('d9c5a4b4-5203-48cc-81b1-9ab42f3a1276', '9ce54ed8-f321-4f48-a492-63e41d287e4d', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('12e049c4-4513-4dd1-9f6e-7f12b925259a', '9ce54ed8-f321-4f48-a492-63e41d287e4d', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),
    ('29f49833-3e29-4721-bec7-4c3db94c02a0', '9ce54ed8-f321-4f48-a492-63e41d287e4d', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('2d8a0b29-f2f7-4a84-a5a5-f0c504fe0b69', '9ce54ed8-f321-4f48-a492-63e41d287e4d', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),

    ('e4df565d-fc9e-404e-8c14-9eef9ac01cb3', 'a67b3065-407c-46b0-b7e8-4e9ba4f71527', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),
    ('ba254071-0eaf-4435-b11e-fc5c2e3c148f', 'a67b3065-407c-46b0-b7e8-4e9ba4f71527', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('e0490a6d-f79b-4c72-a32f-29089e1042f0', 'a67b3065-407c-46b0-b7e8-4e9ba4f71527', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('3d1940ac-8589-43a0-baeb-d7a67d15c159', 'a67b3065-407c-46b0-b7e8-4e9ba4f71527', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('e38a6d32-9293-45b0-8975-ba30de8db081', 'a67b3065-407c-46b0-b7e8-4e9ba4f71527', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),

    ('7b75e75a-0019-47f6-9fd5-6f53d7e36e58', 'd9a7c912-7772-4ecd-aecc-e461a6e0fb97', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),
    ('7294b924-9d3a-4980-841e-ef0d7245b3dc', 'd9a7c912-7772-4ecd-aecc-e461a6e0fb97', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('6a81dd51-8184-4ec1-b8de-09b3bf0340c0', 'd9a7c912-7772-4ecd-aecc-e461a6e0fb97', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('d8a09451-9bb3-497d-931f-04c586d5d37a', 'd9a7c912-7772-4ecd-aecc-e461a6e0fb97', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('6a3308ff-162e-4a9b-b61a-6a31dc43c597', 'd9a7c912-7772-4ecd-aecc-e461a6e0fb97', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f'),

    ('c26d69bc-6186-4454-a14e-ecb7a6e30f57', 'cfc80055-02e3-4e1c-b510-ccf5832a008f', '42dfec60-e6df-4cd6-a1e6-2f228f827b5f'),
    ('0d1263d0-833f-40c0-9560-bd88b6415f6a', 'cfc80055-02e3-4e1c-b510-ccf5832a008f', '7b534913-aeaa-4837-b617-f3e3c504f254'),
    ('d049e680-019e-4b5d-a8b3-9d22506ff6bb', 'cfc80055-02e3-4e1c-b510-ccf5832a008f', 'eb4100eb-b697-43ef-9775-887ab192b0d3'),
    ('00c839c0-f7a7-49b3-b8a7-d16f43f69ac5', 'cfc80055-02e3-4e1c-b510-ccf5832a008f', '616655f0-d066-4ce1-94d5-7d49d95f0bd3'),
    ('2d1dc17b-aa9a-4d4b-b79b-e7b37396471a', 'cfc80055-02e3-4e1c-b510-ccf5832a008f', 'e1e5be66-32ff-4ea7-95e4-3717d725bd7f');


INSERT INTO grade (student_course_id, grade) VALUES
                                                 ('94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', 3.4),
                                                 ('4b8d1eab-845b-4a4f-a6d7-bb30b4a9f8ef', 4.32),
                                                 ('7d29c143-cc48-4c12-b40f-51b3e9f7c04b', 4.2),
                                                 ('e7f87936-1d6b-4bfc-a7d0-686986ef3e6a', 3.95),
                                                 ('f02cb0c2-2758-44d2-b9e7-1d7863d32dcf', 4.1),
                                                 ('29f48565-d833-4753-b0a3-ced061fe717f', 3.8),
                                                 ('afeaa3e3-b9ff-48c4-b324-01e470c7d514', 4.5),
                                                 ('c284c145-75ff-4f6a-b7d9-83d1b9ed601d', 4.73),
                                                 ('4c3bb30d-b4cb-44a1-af46-37c48e46b50e', 3.65),
                                                 ('ac912f08-b3c5-4d9a-a3f8-d51a2a1e9b2b', 4.12),
                                                 ('8c5b3e05-df9d-4d3f-ae1f-2c476ee0d46b', 3.88),
                                                 ('f3f5b9f2-4f1e-4d39-a1a2-16ac8975f3fb', 4.23),
                                                 ('f87c0d5a-f876-4a79-94c6-ba29a1f3bf6f', 4.67),
                                                 ('394cd112-05e0-44cb-a065-cc38b6cb67a2', 3.75),
                                                 ('e1836f70-36b7-4854-b2d6-dce4e470e7cb', 3.2),
                                                 ('4a32b0c5-8d33-4d18-a74a-b46c607c0e39', 4.01),
                                                 ('15b2cb7a-5022-4b82-8a61-55bb0a3d0b9b', 3.92),
                                                 ('f2543035-d27e-47ff-8307-3a51b1d722d8', 3.6),
                                                 ('062f7a7b-8810-4899-b7b6-126a8f5360f2', 4.03),
                                                 ('5534f6a2-d8ec-40d1-8045-8818a4e4d1d8', 4.4),
                                                 ('7ee366eb-13bb-4fc2-a166-af218e90e37a', 3.45),
                                                 ('d9c5a4b4-5203-48cc-81b1-9ab42f3a1276', 3.85),
                                                 ('12e049c4-4513-4dd1-9f6e-7f12b925259a', 4.08),
                                                 ('29f49833-3e29-4721-bec7-4c3db94c02a0', 4.35),
                                                 ('2d8a0b29-f2f7-4a84-a5a5-f0c504fe0b69', 3.55),
                                                 ('e4df565d-fc9e-404e-8c14-9eef9ac01cb3', 4.6),
                                                 ('ba254071-0eaf-4435-b11e-fc5c2e3c148f', 3.72),
                                                 ('e0490a6d-f79b-4c72-a32f-29089e1042f0', 4.28),
                                                 ('3d1940ac-8589-43a0-baeb-d7a67d15c159', 3.25),
                                                 ('e38a6d32-9293-45b0-8975-ba30de8db081', 4.15),
                                                 ('7b75e75a-0019-47f6-9fd5-6f53d7e36e58', 3.68),
                                                 ('7294b924-9d3a-4980-841e-ef0d7245b3dc', 3.82),
                                                 ('6a81dd51-8184-4ec1-b8de-09b3bf0340c0', 4.18),
                                                 ('d8a09451-9bb3-497d-931f-04c586d5d37a', 4.45),
                                                 ('6a3308ff-162e-4a9b-b61a-6a31dc43c597', 3.15),
                                                 ('c26d69bc-6186-4454-a14e-ecb7a6e30f57', 4.07),
                                                 ('0d1263d0-833f-40c0-9560-bd88b6415f6a', 3.98),
                                                 ('d049e680-019e-4b5d-a8b3-9d22506ff6bb', 3.35),
                                                 ('00c839c0-f7a7-49b3-b8a7-d16f43f69ac5', 4.55),
                                                 ('2d1dc17b-aa9a-4d4b-b79b-e7b37396471a', 3.78);
