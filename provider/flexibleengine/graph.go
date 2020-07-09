package flexibleengine

var (
	// nodeTypes all the Nodes that we support right now
	nodeTypes = map[string]struct{}{
		"flexibleengine_antiddos_v1":                        struct{}{},
		"flexibleengine_as_configuration_v1":                struct{}{},
		"flexibleengine_as_group_v1":                        struct{}{},
		"flexibleengine_as_policy_v1":                       struct{}{},
		"flexibleengine_blockstorage_volume_v2":             struct{}{},
		"flexibleengine_cce_cluster_v3":                     struct{}{},
		"flexibleengine_cce_nodes_v3":                       struct{}{},
		"flexibleengine_ces-alarmrule":                      struct{}{},
		"flexibleengine_compute_server_v2":                  struct{}{},
		"flexibleengine_compute_floatingip_associate_v2":    struct{}{},
		"flexibleengine_compute_floatingip_v2":              struct{}{},
		"flexibleengine_compute_instance_v2":                struct{}{},
		"flexibleengine_compute_keypair_v2":                 struct{}{},
		"flexibleengine_compute_servergroup_v2":             struct{}{},
		"flexibleengine_compute_volume_attach_v2":           struct{}{},
		"flexibleengine_csbs_backup_policy_v1":              struct{}{},
		"flexibleengine_csbs_backup_v1":                     struct{}{},
		"flexibleengine_cts_tracker_v1":                     struct{}{},
		"flexibleengine_dcs_instance_v1":                    struct{}{},
		"flexibleengine_dds_instance_v3":                    struct{}{},
		"flexibleengine_dns_recordset_v2":                   struct{}{},
		"flexibleengine_dns_zone_v2":                        struct{}{},
		"flexibleengine_drs_replication_v2":                 struct{}{},
		"flexibleengine_drs_replicationconsistencygroup_v2": struct{}{},
		"flexibleengine_dws_cluster_v1":                     struct{}{},
		"flexibleengine_elb_backend":                        struct{}{},
		"flexibleengine_elb_health":                         struct{}{},
		"flexibleengine_elb_listener":                       struct{}{},
		"flexibleengine_elb_loadbalancer":                   struct{}{},
		"flexibleengine_fw_firewall_group_v2":               struct{}{},
		"flexibleengine_fw_policy_v2":                       struct{}{},
		"flexibleengine_fw_rule_v2":                         struct{}{},
		"flexibleengine_images_image_v2":                    struct{}{},
		"flexibleengine_kms_key_v1":                         struct{}{},
		"flexibleengine_lb_certificate_v2":                  struct{}{},
		"flexibleengine_lb_l7policy_v2":                     struct{}{},
		"flexibleengine_lb_l7rule_v2":                       struct{}{},
		"flexibleengine_lb_loadbalancer_v2":                 struct{}{},
		"flexibleengine_lb_monitor_v2":                      struct{}{},
		"flexibleengine_lb_whitelist_v2":                    struct{}{},
		"flexibleengine_mls_instance_v1":                    struct{}{},
		"flexibleengine_mrs_cluster_v1":                     struct{}{},
		"flexibleengine_mrs_job_v1":                         struct{}{},
		"flexibleengine_nat_dnat_rule_v2":                   struct{}{},
		"flexibleengine_nat_gateway_v2":                     struct{}{},
		"flexibleengine_nat_snat_rule_v2":                   struct{}{},
		"flexibleengine_networking_floatingip_associate_v2": struct{}{},
		"flexibleengine_networking_floatingip_v2":           struct{}{},
		"flexibleengine_networking_network_v2":              struct{}{},
		"flexibleengine_networking_router_interface_v2":     struct{}{},
		"flexibleengine_networking_router_route_v2":         struct{}{},
		"flexibleengine_networking_router_v2":               struct{}{},
		"flexibleengine_networking_subnet_v2":               struct{}{},
		"flexibleengine_networking_vip_associate_v2":        struct{}{},
		"flexibleengine_networking_vip_v2":                  struct{}{},
		"flexibleengine_rds_instance_v1":                    struct{}{},
		"flexibleengine_rds_instance_v3":                    struct{}{},
		"flexibleengine_rds_parametergroup_v3":              struct{}{},
		"flexibleengine_rts_software_config_v1":             struct{}{},
		"flexibleengine_resource_rts_stack_v1":              struct{}{},
		"flexibleengine_s3_bucket":                          struct{}{},
		"flexibleengine_s3_bucket_object":                   struct{}{},
		"flexibleengine_s3_bucket_policy":                   struct{}{},
		"flexibleengine_sfs_file_system_v2":                 struct{}{},
		"flexibleengine_smn_subscription_v2":                struct{}{},
		"flexibleengine_smn_topic_v2":                       struct{}{},
		"flexibleengine-vbs-backup-policy-v2":               struct{}{},
		"flexibleengine-vbs-backup-v2":                      struct{}{},
		"flexibleengine_vpc_eip_v1":                         struct{}{},
		"flexibleengine_vpc_peering_connection_accepter_v2": struct{}{},
		"flexibleengine_vpc_peering_connection_v2":          struct{}{},
		"flexibleengine_vpc_route_v2":                       struct{}{},
		"flexibleengine_vpc_subnet_v1":                      struct{}{},
		"flexibleengine_vpc_v1":                             struct{}{},
	}

	// edgeTypes map of all the supported Edges
	edgeTypes = map[string]struct{}{
		"flexibleengine_compute_interface_attach_v2": struct{}{},
		"flexibleengine_networking_port_v2":          struct{}{},
		"flexibleengine_networking_secgroup_rule_v2": struct{}{},
		"flexibleengine_networking_secgroup_v2":      struct{}{},
		"flexibleengine_lb_listener_v2":              struct{}{},
		"flexibleengine_lb_pool_v2":                  struct{}{},
		"flexibleengine_lb_member_v2":                struct{}{},
	}
)
