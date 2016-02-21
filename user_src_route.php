<?php 
/**
*@author: JJyy
*@todo: get the direct or indirect order transform 
*@param: the slice of some channel_name
*  [1472848505:129:li9t209jm7mc769vmn88o5a345:baidu
*	 " "1472848505:229:li9t27mskkl*c769vmn88o5appp:baidu
*  " "1472848504:239:li9t27mskklc769vmn88o5appp:google]   go slice type
* if get all data start_time=0, end_time=0
* go run main.go ip.go pv.go uv.go user_id.go -uid yahoo 
* 1454035404(start time) 1454035490(end time)
*@return user src route and stay time 
*				 like user_id:src_route:time
*
**/
function user_src_route () {
	// $comm =  "main.exe -ruid ".$channel_name." ".$start_time." ".$end_time;
	$comm =  "main.exe -route_uid 0 0 0";
	//get result
	exec($comm, $res, $rc);	
	// print_r($res);
	//clear the "[ ]"
	$slice = str_replace( array('[', ']'), "", $res[0] );
	$slice_arr = explode(" ", $slice);
	// print_r($slice_arr);
	
	$route_arr = array();	//return route arr
	foreach ($slice_arr as $k => $v) {
		//explede string
		$tmp_arr = explode(':', $v);
		// print_r($tmp_arr);
		if ($tmp_arr[1] != 0) {
			!isset($route_arr[$tmp_arr[1]]) ? $route_arr[$tmp_arr[1]] =
																			$tmp_arr[3].'-' : 
																			$route_arr[$tmp_arr[1]].= 
																			$tmp_arr[3].'-' ;
		}
	}
	
	print_r($route_arr);
  return $route_arr;	
}

user_src_route();

?>


