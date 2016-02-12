<?php 
/**
*@author: JJyy
*@todo: get the direct or indirect order transform 
*@param: the slice of some channel_name  [1472848504:39 1472848504:129 1472848504:139 1472848505:0 1472848505:129 1472848504:39 1472848505:29 1472848504:0 1472848504:129 1472848505:129 1472848505:0 1472848504:229 147284851472848505:139 1472848504:229 1472848504:239 1472848505:229 1472848504:129] go slice type
* if get all data start_time=0, end_time=0
* go run main.go ip.go pv.go uv.go user_id.go -uid yahoo 1454035404(start time) 1454035490(end time)
*@return direct_num:indirect_num
*
**/
// function order_transform( $channel_name, $transform_type, $start_time, $end_time, $db='' ) {
function order_transform( $channel_name, $start_time, $end_time, $db='' ) {
	$comm =  "./go_statis -uid ".$channel_name." ".$start_time." ".$end_time;
	//get result
	exec($comm, $res, $rc);	
	//clear the "[ ]"
	$slice = str_replace( array('[', ']'), "", $res );
	$slice_arr = explode(" ", $slice);
	
	$direct_num = 0;
	$indirect_num = 0;
	
	foreach ($slice_arr as $k => $v) {
		$tmp_arr = explode(":", $v);
		$direct_num += o_direct_tsf( $tmp_arr[1], $tmp_arr[0], $db );
		$indirect_num += o_indirect_tsf( $tmp_arr[1], $channel_name, $tmp_arr[0], $db );
	}
	
	return $direct_tsf+':'+$indirect_tsf;
}

/**
*@author: JJyy
*@todo: get the direct transform number for the user_id 
*		less than two hours
*@param: 
*
**/
function o_direct_tsf($user_id, $timeset, $db='') {
	if ( $user_id == 0 ) {
		return 0;
	}
	else {
		$sql = "select count(order_id) as cn from order_info where order_status=1 and pay_status=2 and confirm_time>".$timeset." and confirm_time<".($timeset + 3600*2);
		$row = $db->query($sql);
		return $row['cn'];
	}
}

/**
*@author: JJyy
*@todo: get the indirect transform number for the user_id 
*		less than 30 days
*@param: 
*
**/
function o_indirect_tsf($user_id, $channel_name, $timeset, $db) {
	$sql = "select count(order_id) as cn from order_info where order_status=1 and pay_status=2 and confirm_time>".$timeset." and confirm_time<".($timeset + 3600*30);
	$row = $db->query($sql);
	$cn = $row['cn'];
	if ( $cn == 0 ) {
		//if within 30 days no order 
		return 0;
	}
	else {
		$rds = new Redis();
		//get the keys for the user_id
		$keys = $rds->keys("*:".$user_id);
		//string like: li9t209jm7mc6m4vmn88o5a7j0:1454035403.8093:10.10.10.29:baidu:0
		$num = 0;
		foreach ($keys as $k => $v) {
			//the time must after timeset
			//same user_id and different channel_name and time after than timeset
			if ( $v[1] > $timeset && $v[3] != $channel_name && $v[4] == $user_id ) {
				$sql = "select count(order_id) as cn from order_info where order_status=1 and pay_status=2 and confirm_time>".$timeset." and confirm_time<".($timeset + 3600*2);
				$row = $db->query($sql);
				$cn = $row['cn'];
				$num += $cn;
			}
		}
		return $num;
	}
}

?>











