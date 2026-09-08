# iPDF

基于 [Wails](https://wails.io/) + Go 的本地 PDF 工具箱，产品形态对标 [PDF24 Tools](https://tools.pdf24.org/zh/all-tools)。

**第一期已实现（全部本地处理，不上传云端）：**

| 工具 | 说明 |
|------|------|
| PDF 合并 | 按顺序合并多个 PDF（pdfcpu） |
| PDF 分割 | 每页 / 页码范围 / 每 N 页（pdfcpu） |
| PDF 压缩 | 结构优化减小体积（pdfcpu optimize） |
| PDF 转图片 | 每页栅格化为 JPG / PNG（go-fitz / MuPDF **静态链接**） |
| 图片转 PDF | JPG / PNG / WEBP → 多页 PDF（pdfcpu） |

## 开发

需要 **CGO + MinGW/GCC**（Windows），以便把 MuPDF 静态库链进可执行文件：

```powershell
cd ipdf
$env:CGO_ENABLED = "1"
wails dev
```

构建：

```powershell
cd ipdf
$env:CGO_ENABLED = "1"
wails build
```

或直接：

```powershell
$env:CGO_ENABLED = "1"
go build -o ipdf.exe .
```

> 默认使用 go-fitz 自带的 MuPDF `.a` 静态库（无需再装系统 MuPDF）。  
> Wails/WebView2 仍依赖系统运行库，这里的「静态」指 **MuPDF 不单独发 DLL**。

## 技术栈

- Wails v2 + Vue 3 + TypeScript + vue-router
- [pdfcpu](https://github.com/pdfcpu/pdfcpu)：合并 / 分割 / 压缩 / 图片转 PDF
- [go-fitz](https://github.com/gen2brain/go-fitz)：PDF 页渲染（CGO 静态链接 MuPDF）

## 限制（第一期）

- 不支持加密 PDF 的解密与处理
- 压缩仅为对象级优化，效果中等（后续可接 Ghostscript）
- 不含 Office / OCR / 水印 / 签署等（首页「即将推出」占位）
- go-fitz 自带库可能不含完整 CJK 字体；复杂中文字体嵌入页请以实测为准
- 所有处理均在本地完成，文件不会上传

## 目录

```
ipdf/
  app.go                    # Wails 绑定
  internal/pdf/             # 工具实现
  internal/fitzstub/        # MinGW 链接 MuPDF 静态库的兼容 stub
  internal/job/             # 临时工作区
  frontend/src/views/       # 工具 UI
```
