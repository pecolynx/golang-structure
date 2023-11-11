create table `article` (
 `id` varchar(26) not null
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp
,`created_by` int not null
,`updated_by` int not null
,`title` varchar(100) not null
,`content` text not null
,primary key(`id`)
,unique(`created_by`, `title`)
,foreign key(`created_by`) references `app_user`(`id`)
,foreign key(`updated_by`) references `app_user`(`id`)
);
