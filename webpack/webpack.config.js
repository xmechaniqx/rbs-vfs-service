const path = require('path')


module.exports = {
  entry: './ui/static/js/main.js', // Точка входа для сборки проекта
mode:'production',
output:{
  path: path.resolve(__dirname,'dist'),
  filename:'index.js'
},
      module: {
    rules: [
      {
        test: /\.css$/, // Регулярное выражение для обработки файлов с расширением .css
        use: ['style-loader', 'css-loader'], // Загрузчики, используемые для обработки CSS-файлов
      },
    ],
    },
devServer:{
  port:9000,
  compress: true,
  hot: true,
  static:{
    directory: path.join(__dirname,'dist')
  }
}
}