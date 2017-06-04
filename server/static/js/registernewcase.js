$(document).ready(function(){
	alert("Hello world")
	$(function(){
		$('.team-list').on("click",".close-div-button",function(){
			if($(this).parents('.team-details').find('input[name^=team_id]').val()!='')
			{
				var id=$(this).parents('.team-details').find('input[name^=team_id]').val();
				var r=confirm('Are you sure you want to delete this experience details?');
				if(r==true)
				{
					$(this).parents('.team-details').remove();
					$.ajax({
						url: "/index.php/creative/delete_team",
						type:"post",
						data:{"id":id},
						success:function(data){
							console.log(data);
						}
					});
					$('.team-list').find('.team-details').last().css("border-bottom","0px");
				}
				else if(r==false){
					return;
				}
			}
			else{

				$(this).parents('.team-details').remove();
				$('.team-list').find('.team-details').last().css("border-bottom","0px");

			}	
		});


		$('.team-list').on("click",".reset-button",function(){
			var team_container=$(this).parents('.team-details');
			console.log(team_container);
			$(team_container).find("input").not("input[name^=team_id]").val('');
			$(team_container).find("textarea").val('');
			$(team_container).find("select").each(function(){
				$('option:selected',this).removeAttr('selected');
			});
			$(team_container).find('.gmaps-output-latitude').html('');
			$(team_container).find('.gmaps-output-longitude').html('');
			$(team_container).find("select[name='from_month[]'] option[value='January']").attr("selected","selected");
			$(team_container).find("select[name='from_year[]'] option[value='2016']").attr("selected","selected");
			$(team_container).find("select[name='to_month[]'] option[value='January']").attr("selected","selected");
			$(team_container).find("select[name='to_year[]'] option[value='2021']").attr("selected","selected");
			$(team_container).find('.to').next().removeClass('inlineBlock');
			$(team_container).find('.to').next().addClass('hidden-div');
			$(team_container).find('.current-position').hide();
			$(team_container).find('.current-field-checkbox').show();
			$(team_container).find('.to-month').show();
			$(team_container).find('.to-year').show();
			$(team_container).find("input[type=checkbox]").prop("checked",false);
			$(team_container).find('.tagit-choice').remove();
		});


		var i=0;var a=0;		
		var country="";
		var state="";
		$('select[name=country]').val(country);
		$('select[name=state]').val(state);
		$('.country').each(function(){
			if($(this).val()!='India')
			{
				$(this).parents('.individual-fields').siblings('.state').hide();
			}
		});
		if($('#location-country').val()!="India")
		{
			$('#location-state').hide();
		}
		else{
			$('select[name=state]').addClass('required');
		}		
	});

	var i=0;
	$('.teams').on("click",function(){
		i++;
		console.log('cp1');
		var team_div	=$('.team-details')[0];
		var cloned_team_div=$(team_div).clone();
		$(cloned_team_div).find("input").val('')
		$(cloned_team_div).find('.tagit-choice,.tagit-new').remove();
		$('.team-details').css("border-bottom","1px solid #DCE0E0");
		$('.team-details').css("margin-bottom","20px");
		$(cloned_team_div).css("border-bottom","0px");
		$(cloned_team_div).find("textarea").html('');
		$(cloned_team_div).find("textarea").val('');
		$(cloned_team_div).find("select").each(function(){
			var value=$(this).val();
			$('option:selected',this).removeAttr('selected');
		});	

		$(cloned_team_div).find('.current-position').removeClass('inlineBlock');
		$(cloned_team_div).find('.current-position').addClass('hidden-div');
		$(cloned_team_div).find("select[name='from_month[]'] option[value='January']").attr("selected","selected");
		$(cloned_team_div).find("select[name='from_year[]'] option[value='2015']").attr("selected","selected");
		$(cloned_team_div).find("select[name='to_month[]'] option[value='January']").attr("selected","selected");
		$(cloned_team_div).find("select[name='to_year[]'] option[value='2015']").attr("selected","selected");
		\
		$(cloned_team_div).find('.tags-div').show();
		$(cloned_team_div).find('.current-position').removeClass('inlineBlock');
		$(cloned_team_div).find('.current-position').addClass('hidden-div');
		$(cloned_team_div).find('.to-month').css("display","inline-block");
		$(cloned_team_div).find('.to-year').css("display","inline-block");
		$(cloned_team_div).find('.close-div-button').show();
		$(cloned_team_div).find('.reset-button').hide();
		$(cloned_team_div).find('input[type=checkbox]').prop("checked",false);
		$(cloned_team_div).find('input[type=checkbox]').val()==0;

		$(cloned_team_div).clone().appendTo('.team-list');

		$('.team-list').on("click","input[type=checkbox]",function(){
			if($(this).prop("checked"))
			{
				console.log("cp1");
				$(this).parents('.experience-checkbox').siblings('.to-month').hide();
				$(this).parents('.experience-checkbox').siblings('.to-year').hide();

				$(this).parents('.experience-checkbox').siblings('.current-position').removeClass('hidden-div');
				$(this).parents('.experience-checkbox').siblings('.current-position').addClass('inlineBlock');
			}
			else
			{
				console.log("cp2");
				$(this).parents('.experience-checkbox').siblings('.current-position').addClass('hidden-div');
				$(this).parents('.experience-checkbox').siblings('.current-position').removeClass('inlineBlock');
				$(this).parents('.experience-checkbox').siblings('.to-month').show();
				$(this).parents('.experience-checkbox').siblings('.to-year').show();
			}
		});
	});
});
