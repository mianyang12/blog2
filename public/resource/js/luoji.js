document.querySelector('#next').onclick = function(){
   let lists =document.querySelectorAll('.item');
   document.querySelector('#slide').appendChild(lists[0]);
}
document.querySelector('#prev').onclick = function(){
  let lists =document.querySelectorAll('.item');
  document.querySelector('#slide').prepend(lists[lists.length-1]);
}
//修改样式较多用className
// const div =document.querySelector('div')
// div.className.'box'
// class是关键字，我们用className
// const div = document.querySelector('div')
// div.className='box'
