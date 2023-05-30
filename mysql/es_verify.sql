CREATE DATABASE if not EXISTS es_verify character set utf8;
create table if not EXISTS  `lb_zone_verify`(
    `id` int(32) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `lb_uuid`  varchar(40) NOT NULL COMMENT 'lb uuid',
    `lb_vm_zone_id` int(11) NOT NULL COMMENT 'zone的id',
    `lb_vm_zone_name`  VARCHAR(100) COMMENT 'zone的name',
    `lb_vm_rs1` VARCHAR(100) COMMENT 'LB rs1的ip',
    `lb_vm_rs2` VARCHAR(12) COMMENT 'LB rs2的ip',
    `lb_vm_secgroupid` varchar(40) NOT NULL COMMENT '安全组',
    `pub_network_id` varchar(40) NOT NULL COMMENT '公网network',
    `pub_subnetwork_id` varchar(40) NOT NULL COMMENT '公网子网subnetwork',
    `pri_network_id` varchar(40) NOT NULL COMMENT '私网network',
    `pri_subnetwork_id` varchar(40) NOT NULL COMMENT '私网子网subnetwork',
    `status`  int(11) NOT NULL COMMENT '验证结果标志位',
    `history_status`  int(11) NOT NULL COMMENT '历史记录标志位',
    `create_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '开始时间',
    `end_time` datetime NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '开始时间',
    PRIMARY KEY(id)
    )ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
