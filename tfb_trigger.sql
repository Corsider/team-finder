PGDMP     *    9            
    {            team_finder2    15.4    15.4 w    s           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            t           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            u           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            v           1262    16785    team_finder2    DATABASE     �   CREATE DATABASE team_finder2 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE team_finder2;
                postgres    false                        2615    2200    public    SCHEMA     2   -- *not* creating schema, since initdb creates it
 2   -- *not* dropping schema, since initdb creates it
                postgres    false            w           0    0    SCHEMA public    ACL     Q   REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;
                   postgres    false    5            �            1255    16941    check_url(text) 	   PROCEDURE     �   CREATE PROCEDURE public.check_url(IN url text)
    LANGUAGE plpgsql
    AS $_$
begin
 IF url !~ '^(http[s]?://)?(www\.)?[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*\.[a-zA-Z]{2,9}(/\S*)?$' then 
  RAISE EXCEPTION 'Not a url';
 END IF;
end; $_$;
 .   DROP PROCEDURE public.check_url(IN url text);
       public          postgres    false    5            �            1255    16942    check_url_fun()    FUNCTION     �   CREATE FUNCTION public.check_url_fun() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
 call check_url(NEW.url);
 RETURN NEW;
end;
$$;
 &   DROP FUNCTION public.check_url_fun();
       public          postgres    false    5            �            1255    16940 +   regist_team(text, text, text, text, bigint)    FUNCTION     `  CREATE FUNCTION public.regist_team(name text, description text, rules text, place text, creator_id bigint) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
declare 
  new_id bigint;
begin
 if creator_id is NULL or creator_id < 0 THEN
  RAISE EXCEPTION 'Not valid creator id';
 END IF;
 INSERT INTO team (name, rate, description, rules, reg_date, place) VALUES (name, 0.0, description, rules, CURRENT_TIMESTAMP, place) RETURNING team_id into new_id;
 insert into user_team (team_id, user_id, role, date_of_entry, hidden) values (new_id, creator_id, 'Creator', CURRENT_TIMESTAMP, false);
 return new_id;
end;
$$;
 j   DROP FUNCTION public.regist_team(name text, description text, rules text, place text, creator_id bigint);
       public          postgres    false    5            �            1259    16786    events    TABLE     �   CREATE TABLE public.events (
    event_id bigint NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    date date NOT NULL,
    online boolean NOT NULL,
    main_theme text NOT NULL,
    url text NOT NULL,
    creator_id bigint NOT NULL
);
    DROP TABLE public.events;
       public         heap    postgres    false    5            �            1259    16791    events_creator_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.events_creator_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.events_creator_id_seq;
       public          postgres    false    214    5            x           0    0    events_creator_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.events_creator_id_seq OWNED BY public.events.creator_id;
          public          postgres    false    215            �            1259    16792    events_event_id_seq    SEQUENCE     |   CREATE SEQUENCE public.events_event_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.events_event_id_seq;
       public          postgres    false    5    214            y           0    0    events_event_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.events_event_id_seq OWNED BY public.events.event_id;
          public          postgres    false    216            �            1259    16793    events_tags    TABLE     `   CREATE TABLE public.events_tags (
    event_id integer NOT NULL,
    tag_id integer NOT NULL
);
    DROP TABLE public.events_tags;
       public         heap    postgres    false    5            �            1259    16796    events_tags_event_id_seq    SEQUENCE     �   CREATE SEQUENCE public.events_tags_event_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.events_tags_event_id_seq;
       public          postgres    false    5    217            z           0    0    events_tags_event_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.events_tags_event_id_seq OWNED BY public.events_tags.event_id;
          public          postgres    false    218            �            1259    16797    events_tags_tag_id_seq    SEQUENCE     �   CREATE SEQUENCE public.events_tags_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.events_tags_tag_id_seq;
       public          postgres    false    217    5            {           0    0    events_tags_tag_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.events_tags_tag_id_seq OWNED BY public.events_tags.tag_id;
          public          postgres    false    219            �            1259    16798 
   global_tag    TABLE     b   CREATE TABLE public.global_tag (
    global_tag_id bigint NOT NULL,
    category text NOT NULL
);
    DROP TABLE public.global_tag;
       public         heap    postgres    false    5            �            1259    16803    global_tag_global_tag_id_seq    SEQUENCE     �   CREATE SEQUENCE public.global_tag_global_tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 3   DROP SEQUENCE public.global_tag_global_tag_id_seq;
       public          postgres    false    220    5            |           0    0    global_tag_global_tag_id_seq    SEQUENCE OWNED BY     ]   ALTER SEQUENCE public.global_tag_global_tag_id_seq OWNED BY public.global_tag.global_tag_id;
          public          postgres    false    221            �            1259    16804    tag    TABLE     w   CREATE TABLE public.tag (
    tag_id bigint NOT NULL,
    activity text NOT NULL,
    global_tag_id bigint NOT NULL
);
    DROP TABLE public.tag;
       public         heap    postgres    false    5            �            1259    16809    tag_global_tag_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.tag_global_tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.tag_global_tag_id_seq;
       public          postgres    false    222    5            }           0    0    tag_global_tag_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.tag_global_tag_id_seq OWNED BY public.tag.global_tag_id;
          public          postgres    false    223            �            1259    16810    tag_tag_id_seq    SEQUENCE     w   CREATE SEQUENCE public.tag_tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.tag_tag_id_seq;
       public          postgres    false    222    5            ~           0    0    tag_tag_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.tag_tag_id_seq OWNED BY public.tag.tag_id;
          public          postgres    false    224            �            1259    16811    team    TABLE     �   CREATE TABLE public.team (
    team_id bigint NOT NULL,
    name text,
    rate double precision,
    description text,
    rules text,
    reg_date date,
    place text
);
    DROP TABLE public.team;
       public         heap    postgres    false    5            �            1259    16816 
   team_event    TABLE     |   CREATE TABLE public.team_event (
    event_id integer NOT NULL,
    team_id integer NOT NULL,
    reg_time date NOT NULL
);
    DROP TABLE public.team_event;
       public         heap    postgres    false    5            �            1259    16819    team_event_event_id_seq    SEQUENCE     �   CREATE SEQUENCE public.team_event_event_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.team_event_event_id_seq;
       public          postgres    false    5    226                       0    0    team_event_event_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.team_event_event_id_seq OWNED BY public.team_event.event_id;
          public          postgres    false    227            �            1259    16820    team_event_team_id_seq    SEQUENCE     �   CREATE SEQUENCE public.team_event_team_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.team_event_team_id_seq;
       public          postgres    false    5    226            �           0    0    team_event_team_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.team_event_team_id_seq OWNED BY public.team_event.team_id;
          public          postgres    false    228            �            1259    16821 	   team_tags    TABLE     [   CREATE TABLE public.team_tags (
    tag_id bigint NOT NULL,
    team_id bigint NOT NULL
);
    DROP TABLE public.team_tags;
       public         heap    postgres    false    5            �            1259    16824    team_tags_tag_id_seq    SEQUENCE     }   CREATE SEQUENCE public.team_tags_tag_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.team_tags_tag_id_seq;
       public          postgres    false    5    229            �           0    0    team_tags_tag_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.team_tags_tag_id_seq OWNED BY public.team_tags.tag_id;
          public          postgres    false    230            �            1259    16825    team_tags_team_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.team_tags_team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.team_tags_team_id_seq;
       public          postgres    false    229    5            �           0    0    team_tags_team_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.team_tags_team_id_seq OWNED BY public.team_tags.team_id;
          public          postgres    false    231            �            1259    16826    team_team_id_seq    SEQUENCE     y   CREATE SEQUENCE public.team_team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.team_team_id_seq;
       public          postgres    false    5    225            �           0    0    team_team_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.team_team_id_seq OWNED BY public.team.team_id;
          public          postgres    false    232            �            1259    16827 	   user_team    TABLE     �   CREATE TABLE public.user_team (
    team_id bigint NOT NULL,
    user_id bigint NOT NULL,
    role text NOT NULL,
    date_of_entry date NOT NULL,
    hidden boolean NOT NULL
);
    DROP TABLE public.user_team;
       public         heap    postgres    false    5            �            1259    16832    user_team_team_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.user_team_team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.user_team_team_id_seq;
       public          postgres    false    233    5            �           0    0    user_team_team_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.user_team_team_id_seq OWNED BY public.user_team.team_id;
          public          postgres    false    234            �            1259    16833    user_team_user_id_seq    SEQUENCE     ~   CREATE SEQUENCE public.user_team_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.user_team_user_id_seq;
       public          postgres    false    5    233            �           0    0    user_team_user_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.user_team_user_id_seq OWNED BY public.user_team.user_id;
          public          postgres    false    235            �            1259    16834    users    TABLE     �   CREATE TABLE public.users (
    user_id bigint NOT NULL,
    name text,
    nickname text,
    rate numeric(10,8),
    description text NOT NULL,
    login text,
    password text
);
    DROP TABLE public.users;
       public         heap    postgres    false    5            �            1259    16839 
   users_tags    TABLE     ^   CREATE TABLE public.users_tags (
    tag_id integer NOT NULL,
    user_id integer NOT NULL
);
    DROP TABLE public.users_tags;
       public         heap    postgres    false    5            �            1259    16842    users_tags_tag_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_tags_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.users_tags_tag_id_seq;
       public          postgres    false    5    237            �           0    0    users_tags_tag_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.users_tags_tag_id_seq OWNED BY public.users_tags.tag_id;
          public          postgres    false    238            �            1259    16843    users_tags_user_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_tags_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.users_tags_user_id_seq;
       public          postgres    false    5    237            �           0    0    users_tags_user_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.users_tags_user_id_seq OWNED BY public.users_tags.user_id;
          public          postgres    false    239            �            1259    16844    users_user_id_seq    SEQUENCE     z   CREATE SEQUENCE public.users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_user_id_seq;
       public          postgres    false    5    236            �           0    0    users_user_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;
          public          postgres    false    240            �           2604    16845    events event_id    DEFAULT     r   ALTER TABLE ONLY public.events ALTER COLUMN event_id SET DEFAULT nextval('public.events_event_id_seq'::regclass);
 >   ALTER TABLE public.events ALTER COLUMN event_id DROP DEFAULT;
       public          postgres    false    216    214            �           2604    16846    events creator_id    DEFAULT     v   ALTER TABLE ONLY public.events ALTER COLUMN creator_id SET DEFAULT nextval('public.events_creator_id_seq'::regclass);
 @   ALTER TABLE public.events ALTER COLUMN creator_id DROP DEFAULT;
       public          postgres    false    215    214            �           2604    16847    events_tags event_id    DEFAULT     |   ALTER TABLE ONLY public.events_tags ALTER COLUMN event_id SET DEFAULT nextval('public.events_tags_event_id_seq'::regclass);
 C   ALTER TABLE public.events_tags ALTER COLUMN event_id DROP DEFAULT;
       public          postgres    false    218    217            �           2604    16848    events_tags tag_id    DEFAULT     x   ALTER TABLE ONLY public.events_tags ALTER COLUMN tag_id SET DEFAULT nextval('public.events_tags_tag_id_seq'::regclass);
 A   ALTER TABLE public.events_tags ALTER COLUMN tag_id DROP DEFAULT;
       public          postgres    false    219    217            �           2604    16849    global_tag global_tag_id    DEFAULT     �   ALTER TABLE ONLY public.global_tag ALTER COLUMN global_tag_id SET DEFAULT nextval('public.global_tag_global_tag_id_seq'::regclass);
 G   ALTER TABLE public.global_tag ALTER COLUMN global_tag_id DROP DEFAULT;
       public          postgres    false    221    220            �           2604    16850 
   tag tag_id    DEFAULT     h   ALTER TABLE ONLY public.tag ALTER COLUMN tag_id SET DEFAULT nextval('public.tag_tag_id_seq'::regclass);
 9   ALTER TABLE public.tag ALTER COLUMN tag_id DROP DEFAULT;
       public          postgres    false    224    222            �           2604    16851    tag global_tag_id    DEFAULT     v   ALTER TABLE ONLY public.tag ALTER COLUMN global_tag_id SET DEFAULT nextval('public.tag_global_tag_id_seq'::regclass);
 @   ALTER TABLE public.tag ALTER COLUMN global_tag_id DROP DEFAULT;
       public          postgres    false    223    222            �           2604    16852    team team_id    DEFAULT     l   ALTER TABLE ONLY public.team ALTER COLUMN team_id SET DEFAULT nextval('public.team_team_id_seq'::regclass);
 ;   ALTER TABLE public.team ALTER COLUMN team_id DROP DEFAULT;
       public          postgres    false    232    225            �           2604    16853    team_event event_id    DEFAULT     z   ALTER TABLE ONLY public.team_event ALTER COLUMN event_id SET DEFAULT nextval('public.team_event_event_id_seq'::regclass);
 B   ALTER TABLE public.team_event ALTER COLUMN event_id DROP DEFAULT;
       public          postgres    false    227    226            �           2604    16854    team_event team_id    DEFAULT     x   ALTER TABLE ONLY public.team_event ALTER COLUMN team_id SET DEFAULT nextval('public.team_event_team_id_seq'::regclass);
 A   ALTER TABLE public.team_event ALTER COLUMN team_id DROP DEFAULT;
       public          postgres    false    228    226            �           2604    16855    team_tags tag_id    DEFAULT     t   ALTER TABLE ONLY public.team_tags ALTER COLUMN tag_id SET DEFAULT nextval('public.team_tags_tag_id_seq'::regclass);
 ?   ALTER TABLE public.team_tags ALTER COLUMN tag_id DROP DEFAULT;
       public          postgres    false    230    229            �           2604    16856    team_tags team_id    DEFAULT     v   ALTER TABLE ONLY public.team_tags ALTER COLUMN team_id SET DEFAULT nextval('public.team_tags_team_id_seq'::regclass);
 @   ALTER TABLE public.team_tags ALTER COLUMN team_id DROP DEFAULT;
       public          postgres    false    231    229            �           2604    16857    user_team team_id    DEFAULT     v   ALTER TABLE ONLY public.user_team ALTER COLUMN team_id SET DEFAULT nextval('public.user_team_team_id_seq'::regclass);
 @   ALTER TABLE public.user_team ALTER COLUMN team_id DROP DEFAULT;
       public          postgres    false    234    233            �           2604    16858    user_team user_id    DEFAULT     v   ALTER TABLE ONLY public.user_team ALTER COLUMN user_id SET DEFAULT nextval('public.user_team_user_id_seq'::regclass);
 @   ALTER TABLE public.user_team ALTER COLUMN user_id DROP DEFAULT;
       public          postgres    false    235    233            �           2604    16859    users user_id    DEFAULT     n   ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
       public          postgres    false    240    236            �           2604    16860    users_tags tag_id    DEFAULT     v   ALTER TABLE ONLY public.users_tags ALTER COLUMN tag_id SET DEFAULT nextval('public.users_tags_tag_id_seq'::regclass);
 @   ALTER TABLE public.users_tags ALTER COLUMN tag_id DROP DEFAULT;
       public          postgres    false    238    237            �           2604    16861    users_tags user_id    DEFAULT     x   ALTER TABLE ONLY public.users_tags ALTER COLUMN user_id SET DEFAULT nextval('public.users_tags_user_id_seq'::regclass);
 A   ALTER TABLE public.users_tags ALTER COLUMN user_id DROP DEFAULT;
       public          postgres    false    239    237            V          0    16786    events 
   TABLE DATA           h   COPY public.events (event_id, name, description, date, online, main_theme, url, creator_id) FROM stdin;
    public          postgres    false    214   [�       Y          0    16793    events_tags 
   TABLE DATA           7   COPY public.events_tags (event_id, tag_id) FROM stdin;
    public          postgres    false    217   &�       \          0    16798 
   global_tag 
   TABLE DATA           =   COPY public.global_tag (global_tag_id, category) FROM stdin;
    public          postgres    false    220   M�       ^          0    16804    tag 
   TABLE DATA           >   COPY public.tag (tag_id, activity, global_tag_id) FROM stdin;
    public          postgres    false    222   �       a          0    16811    team 
   TABLE DATA           X   COPY public.team (team_id, name, rate, description, rules, reg_date, place) FROM stdin;
    public          postgres    false    225   �       b          0    16816 
   team_event 
   TABLE DATA           A   COPY public.team_event (event_id, team_id, reg_time) FROM stdin;
    public          postgres    false    226   ��       e          0    16821 	   team_tags 
   TABLE DATA           4   COPY public.team_tags (tag_id, team_id) FROM stdin;
    public          postgres    false    229   ݊       i          0    16827 	   user_team 
   TABLE DATA           R   COPY public.user_team (team_id, user_id, role, date_of_entry, hidden) FROM stdin;
    public          postgres    false    233   �       l          0    16834    users 
   TABLE DATA           \   COPY public.users (user_id, name, nickname, rate, description, login, password) FROM stdin;
    public          postgres    false    236   M�       m          0    16839 
   users_tags 
   TABLE DATA           5   COPY public.users_tags (tag_id, user_id) FROM stdin;
    public          postgres    false    237   �       �           0    0    events_creator_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.events_creator_id_seq', 1, false);
          public          postgres    false    215            �           0    0    events_event_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.events_event_id_seq', 6, true);
          public          postgres    false    216            �           0    0    events_tags_event_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.events_tags_event_id_seq', 1, false);
          public          postgres    false    218            �           0    0    events_tags_tag_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.events_tags_tag_id_seq', 1, false);
          public          postgres    false    219            �           0    0    global_tag_global_tag_id_seq    SEQUENCE SET     K   SELECT pg_catalog.setval('public.global_tag_global_tag_id_seq', 12, true);
          public          postgres    false    221            �           0    0    tag_global_tag_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.tag_global_tag_id_seq', 1, false);
          public          postgres    false    223            �           0    0    tag_tag_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.tag_tag_id_seq', 23, true);
          public          postgres    false    224            �           0    0    team_event_event_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.team_event_event_id_seq', 1, false);
          public          postgres    false    227            �           0    0    team_event_team_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.team_event_team_id_seq', 1, false);
          public          postgres    false    228            �           0    0    team_tags_tag_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.team_tags_tag_id_seq', 1, false);
          public          postgres    false    230            �           0    0    team_tags_team_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.team_tags_team_id_seq', 1, false);
          public          postgres    false    231            �           0    0    team_team_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.team_team_id_seq', 9, true);
          public          postgres    false    232            �           0    0    user_team_team_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.user_team_team_id_seq', 1, false);
          public          postgres    false    234            �           0    0    user_team_user_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.user_team_user_id_seq', 1, false);
          public          postgres    false    235            �           0    0    users_tags_tag_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.users_tags_tag_id_seq', 1, false);
          public          postgres    false    238            �           0    0    users_tags_user_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.users_tags_user_id_seq', 1, false);
          public          postgres    false    239            �           0    0    users_user_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.users_user_id_seq', 12, true);
          public          postgres    false    240            �           2606    16863    events events_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (event_id);
 <   ALTER TABLE ONLY public.events DROP CONSTRAINT events_pkey;
       public            postgres    false    214            �           2606    16865    global_tag global_tag_pkey 
   CONSTRAINT     c   ALTER TABLE ONLY public.global_tag
    ADD CONSTRAINT global_tag_pkey PRIMARY KEY (global_tag_id);
 D   ALTER TABLE ONLY public.global_tag DROP CONSTRAINT global_tag_pkey;
       public            postgres    false    220            �           2606    16867    tag tag_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (tag_id);
 6   ALTER TABLE ONLY public.tag DROP CONSTRAINT tag_pkey;
       public            postgres    false    222            �           2606    16869    team team_pkey 
   CONSTRAINT     Q   ALTER TABLE ONLY public.team
    ADD CONSTRAINT team_pkey PRIMARY KEY (team_id);
 8   ALTER TABLE ONLY public.team DROP CONSTRAINT team_pkey;
       public            postgres    false    225            �           2606    16871    users users_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    236            �           1259    16872    fki_g    INDEX     >   CREATE INDEX fki_g ON public.team_tags USING btree (team_id);
    DROP INDEX public.fki_g;
       public            postgres    false    229            �           1259    16873 
   fki_tag_FK    INDEX     D   CREATE INDEX "fki_tag_FK" ON public.team_tags USING btree (tag_id);
     DROP INDEX public."fki_tag_FK";
       public            postgres    false    229            �           1259    16874    team_to_tag_IDX    INDEX     �   CREATE UNIQUE INDEX "team_to_tag_IDX" ON public.team_tags USING btree (tag_id, team_id);

ALTER TABLE public.team_tags CLUSTER ON "team_to_tag_IDX";
 %   DROP INDEX public."team_to_tag_IDX";
       public            postgres    false    229    229            �           1259    16875    user_team_idx    INDEX     �   CREATE UNIQUE INDEX user_team_idx ON public.user_team USING btree (team_id, user_id);

ALTER TABLE public.user_team CLUSTER ON user_team_idx;
 !   DROP INDEX public.user_team_idx;
       public            postgres    false    233    233            �           2620    16943    events check_event    TRIGGER     z   CREATE TRIGGER check_event BEFORE INSERT OR UPDATE ON public.events FOR EACH ROW EXECUTE FUNCTION public.check_url_fun();
 +   DROP TRIGGER check_event ON public.events;
       public          postgres    false    214    243            �           2606    16876    events_tags event_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.events_tags
    ADD CONSTRAINT event_id FOREIGN KEY (event_id) REFERENCES public.events(event_id) MATCH FULL;
 >   ALTER TABLE ONLY public.events_tags DROP CONSTRAINT event_id;
       public          postgres    false    3246    217    214            �           2606    16881    team_event event_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.team_event
    ADD CONSTRAINT event_id FOREIGN KEY (event_id) REFERENCES public.events(event_id) MATCH FULL;
 =   ALTER TABLE ONLY public.team_event DROP CONSTRAINT event_id;
       public          postgres    false    214    3246    226            �           2606    16886    tag global_tag_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.tag
    ADD CONSTRAINT global_tag_foreign FOREIGN KEY (global_tag_id) REFERENCES public.global_tag(global_tag_id) MATCH FULL;
 @   ALTER TABLE ONLY public.tag DROP CONSTRAINT global_tag_foreign;
       public          postgres    false    222    220    3248            �           2606    16891    team_tags tag_FK    FK CONSTRAINT     �   ALTER TABLE ONLY public.team_tags
    ADD CONSTRAINT "tag_FK" FOREIGN KEY (tag_id) REFERENCES public.tag(tag_id) MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
 <   ALTER TABLE ONLY public.team_tags DROP CONSTRAINT "tag_FK";
       public          postgres    false    3250    229    222            �           2606    16896    users_tags tag_id    FK CONSTRAINT     |   ALTER TABLE ONLY public.users_tags
    ADD CONSTRAINT tag_id FOREIGN KEY (tag_id) REFERENCES public.tag(tag_id) MATCH FULL;
 ;   ALTER TABLE ONLY public.users_tags DROP CONSTRAINT tag_id;
       public          postgres    false    237    222    3250            �           2606    16901    events_tags tag_id    FK CONSTRAINT     }   ALTER TABLE ONLY public.events_tags
    ADD CONSTRAINT tag_id FOREIGN KEY (tag_id) REFERENCES public.tag(tag_id) MATCH FULL;
 <   ALTER TABLE ONLY public.events_tags DROP CONSTRAINT tag_id;
       public          postgres    false    217    3250    222            �           2606    16906    team_tags team_FK    FK CONSTRAINT     �   ALTER TABLE ONLY public.team_tags
    ADD CONSTRAINT "team_FK" FOREIGN KEY (team_id) REFERENCES public.team(team_id) MATCH FULL ON UPDATE CASCADE ON DELETE CASCADE;
 =   ALTER TABLE ONLY public.team_tags DROP CONSTRAINT "team_FK";
       public          postgres    false    225    229    3252            �           2606    16911    user_team team_constaint    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_team
    ADD CONSTRAINT team_constaint FOREIGN KEY (team_id) REFERENCES public.team(team_id) MATCH FULL;
 B   ALTER TABLE ONLY public.user_team DROP CONSTRAINT team_constaint;
       public          postgres    false    225    233    3252            �           2606    16916    team_event team_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.team_event
    ADD CONSTRAINT team_id FOREIGN KEY (team_id) REFERENCES public.team(team_id) MATCH FULL;
 <   ALTER TABLE ONLY public.team_event DROP CONSTRAINT team_id;
       public          postgres    false    225    3252    226            �           2606    16921    events user_FK    FK CONSTRAINT     �   ALTER TABLE ONLY public.events
    ADD CONSTRAINT "user_FK" FOREIGN KEY (creator_id) REFERENCES public.users(user_id) MATCH FULL;
 :   ALTER TABLE ONLY public.events DROP CONSTRAINT "user_FK";
       public          postgres    false    214    236    3258            �           2606    16926    user_team user_constraint    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_team
    ADD CONSTRAINT user_constraint FOREIGN KEY (user_id) REFERENCES public.users(user_id) MATCH FULL;
 C   ALTER TABLE ONLY public.user_team DROP CONSTRAINT user_constraint;
       public          postgres    false    236    3258    233            �           2606    16931    users_tags user_id    FK CONSTRAINT     �   ALTER TABLE ONLY public.users_tags
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.users(user_id) MATCH FULL;
 <   ALTER TABLE ONLY public.users_tags DROP CONSTRAINT user_id;
       public          postgres    false    236    3258    237            V   �   x�=�I
�@E�է�(I���&�*8�gI@<�7hc�!g�u#�0���׫�;$�pG�	j>��x��U��x�AA����VЊO����2�J��q�'�^':�bjE��8�㽰��=q*�]|�j8�Ean�=��S���*�/p}-��
��Ź0�^�xp�E��d���|�]�əa��>0��      Y      x�3�44�2�4������ ��      \   �   x�=�M�@���)8�	*x��`Hĸ`�$����|� �N���M5"z+�ᇕ��Y�+>�;�|q��BCЋ���r�+�ԫ�}�-���<����I�v0Y�A��~.b�&x����#��T���h�      ^   &  x�EO�JQ]����7���'�2�ڴ�&0�l�(Ea�hSP�>`2��t��?�<�\<ν�s�Z�:Wq�+�B�X��vy��d;
��EȺ"���nK���I'v���(�/�|i��Ě��u�2��H��tI���1�r������ILB&��1c��=Vt9�����#��e���9h�a=�=��lW�#�E}�Wm��\ciFt����/���-���tM��	���ǭ9VV�KRUn����z����xn�����e�g2�b�LS�2�������KXM��E��/�ڮo      a   �   x�3�,IM���K�4�,OUH,JU�+�MJ-R��%�*��d��r��p^�sa����.l���ˈ37�$C!'�,����X�d
�� ��/U�--.A��20�"CK�Q;.6p��qqq d/�      b   $   x�3�4�4202�5"#.#NC�̵������ c�t      e      x�34�4�2��4����� O�      i   9   x�3�4��IML�4202�54�54�L�2�4��M�MJ-����Zr�p��qqq �
c      l   �   x�3���,����LJ��4ѳ4� μ���̼t��ԢT�l���������\Δ�t��toosc�xc#.#N���"NCCC�f� ��XlfV��
3���%5/����q�așR�d
P���">�#8�+F��� ��<�      m   *   x�3�4�24 @ӈ��́ؐӘ��D��=... ���     