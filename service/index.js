const http = require('http')
const fs = require('fs')
const path = require('path')

http.createServer((req, res) => {
  const fileName = path.basename(req.url)
  fileType = path.extname(fileName)
  filePath = fileName ? `${__dirname}/${fileName}` : `${__dirname}/index.html`
  if (!fs.existsSync(filePath)) res.end('')
  res.end(fs.readFileSync(filePath))
}).listen(8080, '0.0.0.0')