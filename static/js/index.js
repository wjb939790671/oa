$(function () {
  
    loadMainMenu();
    
   
})


function loadMainMenu() {
    $('.easyui-tree').tree({
        onClick: function (node) {
            if ($('.easyui-tree').tree('isLeaf', node.target)) {            
                if ($('#tabs').tabs('exists', node.text)) {
                    $('#tabs').tabs('select',  node.text);                 
                }else{
                    addTab(node.text,node.attributes.url);
                }
                
            } else {
                $(this).tree(node.state === 'closed' ? 'expand' : 'collapse', node.target);
            }
    
    
            // alert node text property when clicked
        }
    });
}

function addTab(title, url) {   
    var iframe ='<iframe src="'+url+'" frameborder="0" style="height: 100%;width: 100%;backgroud-color:red;" scrolling="yes"></iframe>';
        $('#tabs').tabs('add', {
                title: title,
                closable: true,
                content:iframe
            });
   
}



