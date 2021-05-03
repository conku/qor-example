function materialSwitch(url){
    //var input = $("#switch_"+id)
    $.ajax({
        url: url,
        type: 'PUT',
        success: function() {
            //console.log(response)
        }
    });
    // console.log(id)
    // console.log(input)
    // console.log(input.prop("checked"))
}


function storeState(ids){
    //pdd_stores_2_state
    console.log(ids)
    $.ajax({
        url: 'script.php',
        type: 'PUT',
        success: function( response ) {
        }
     });
}