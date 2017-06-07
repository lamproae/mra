var prec=1;
$("#precondition_add_assert").click(function(){
	$('#preconditionassert_'+prec).html('	<td class="text-center">'+prec+'</td>											<td class="text-center"><input type="text" name="preconditionassertdut_'+prec+'"  placeholder="enable" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="preconditionassertmode_'+prec+'"  placeholder="enable" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="preconditionassertcli_'+prec+'" placeholder="show runngin-config" class="form-control col-sm-4"/></td>											<td class="text-center"><input type="text" name="preconditionassertexpected_'+prec+'" placeholder="br1000[[_space_]]+up" class="form-control col-sm-4"/></td>');

	$('#preconditionassert_tab').append('<tr id="preconditionassert_'+(prec+1)+'"></tr>');
	prec++; 
});
$("#precondition_delete_assert").click(function(){
	if(prec>1){
		$("#preconditionassert_"+(prec-1)).html('');
		prec--;
	}
});

var posc=1;
$("#postcondition_add_assert").click(function(){
	$('#postconditionassert_'+posc).html('	<td class="text-center">'+posc+'</td>											<td class="text-center"><input type="text" name="postconditionassertdut_'+posc+'"  placeholder="enable" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="postconditionassertmode_'+posc+'"  placeholder="enable" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="postconditionassertcli_'+posc+'" placeholder="show runngin-config" class="form-control col-sm-4"/></td>											<td class="text-center"><input type="text" name="postconditionassertexpected_'+posc+'" placeholder="br1000[[_space_]]+up" class="form-control col-sm-4"/></td>');

	$('#postconditionassert_tab').append('<tr id="postconditionassert_'+(posc+1)+'"></tr>');
	posc++; 
});
$("#postcondition_delete_assert").click(function(){
	if(posc>1){
		$("#postconditionassert_"+(posc-1)).html('');
		posc--;
	}
});


var comc=1;
$("#routine_add_command").click(function(){
	$('#routine_command_'+comc).html('	<td class="text-center">'+comc+'</td>											<td class="text-center"><input type="text" name="routine_command_'+comc+'_dut"  placeholder="DUT1" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="routine_command_'+comc+'_mode"  placeholder="enable" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="routine_command_'+comc+'_cli" placeholder="show runngin-config" class="form-control col-sm-4"/></td>											<td class="text-center"><input type="text" name="routine_command_'+comc+'_expected" placeholder="br1000[[_space_]]+up" class="form-control col-sm-4"/></td>');

	$('#routine_command_tab').append('<tr id="routine_command_'+(comc+1)+'"></tr>');
	comc++; 
});
$("#routine_delete_command").click(function(){
	if(comc>1){
		$("#routine_command_"+(comc-1)).html('');
		comc--;
	}
});

var clcomc=1;
$("#clear_add_command").click(function(){
	$('#clear_command_'+clcomc).html('	<td class="text-center">'+clcomc+'</td>											<td class="text-center"><input type="text" name="clear_command_'+clcomc+'_dut"  placeholder="DUT1" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="clear_command_'+clcomc+'_mode"  placeholder="enable" class="form-control col-sm-1"/></td>											<td class="text-center"><input type="text" name="clear_command_'+clcomc+'_cli" placeholder="show runngin-config" class="form-control col-sm-4"/></td>											<td class="text-center"><input type="text" name="clear_command_'+clcomc+'_expected" placeholder="br1000[[_space_]]+up" class="form-control col-sm-4"/></td>');

	$('#clear_command_tab').append('<tr id="clear_command_'+(clcomc+1)+'"></tr>');
	clcomc++; 
});

$("#clear_delete_command").click(function(){
	if(clcomc>1){
		$("#clear_command_"+(clcomc-1)).html('');
		clcomc--;
	}
});
