<?php 
/**
*@author: JJyy
*@todo: get the direct or indirect register transform 
*@param: the slice of some channel_name  [1472848504:39 1472848504:129 1472848504:139 1472848505:0 1472848505:129 1472848504:39 1472848505:29 1472848504:0 1472848504:129 1472848505:129 1472848505:0 1472848504:229 147284851472848505:139 1472848504:229 1472848504:239 1472848505:229 1472848504:129] go slice type
* if get all data start_time=0, end_time=0
* go run main.go ip.go pv.go uv.go user_id.go -uid yahoo 1454035404(start time) 1454035490(end time)
*@return direct_num:indirect_num
*
**/
// function order_transform( $channel_name, $transform_type, $start_time, $end_time, $db='' ) {
function reg_transform( $channel_name, $start_time, $end_time, $db='' ) {
	$comm =  "./go_statis -ruid ".$channel_name." ".$start_time." ".$end_time;
	//get result
	exec($comm, $res, $rc);	
	//clear the "[ ]"
	$slice = str_replace( array('[', ']'), "", $res );
	$slice_arr = explode(" ", $slice);
	
	$real_user_ids = array(); 
	$sessid_arr = array(); 
	foreach ($slice_arr as $k => $v) {
		$tmp_arr = explode(":", $v);
		//if user_id != 0, is direct register user
		if( $tmp_arr[1] != 0 ) {
			$real_user_ids[] = $tmp_arr[1];
		}
		else if ( $tmp_arr[1] == 0 ) {	//if user_id == 0, then check within 30 days reg user
			$sessid_arr[] = $tmp_arr[2];
		}
	}
	
	//del the repeat user_id
	$uni_user_ids = array_unique($real_user_ids);
	$direct_reg_num = count($uni_user_ids);		//direct register user count
	
	$indirect_reg_num = r_indirect_tsf($sessid_arr);
	
	return $direct_reg_num.':'.$indirect_reg_num;
}


/**
*@author: JJyy
*@todo: get the indirect user_id  
*@param: 
*@return: 
*
**/
function r_indirect_tsf($sessid_arr) {
	$rds = new Redis();
	$tmp = array();
	$num = 0;
	/**
	* if get the keys is more than two, just check the second key user_id value
	* if the value is 0, no reg, if not, registed
	* every channel_name for every client( sessid ) just have one indirect_reg 
	* at most
	**/
	foreach ($sessid_arr as $k => $v) {
		$tmp[] = $rds->keys( $v."*" );	//two-dimemsion array
	}

	foreach ($tmp as $k => $v) {
		
		if ( count($v) > 1 ) {
			//get the check time
			$tmp_check_time = explode(':', $v[0]);
			$check_time = $tmp_check_time[1];
			
			$indirect_key = $v[1];	//the second key is indirect_key
			//key like: li9t209jm7mc6m4vmn88o5a7j0:1454035403.8093:10.10.10.29:baidu:0
			$key_arr = explode(':', $indirect_key);
			if ( $key_arr[4] != 0 && $key_arr[1] < ($check_time + 3600*30)  ) {	//the user_id !=0 and the time is less than 30 days
				$num += 1;
			}
		}
	} //end foreach
	
	return $num;
}



