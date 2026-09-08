# iPDF

基于 [Wails](https://wails.io/) + Go 的本地 PDF 工具箱，产品形态对标 [PDF24 Tools](https://tools.pdf24.org/zh/all-tools)。

**第一期已实现（全部本地处理，不上传云端）：**

| 工具 | 说明 |
|------|------|
| PDF 合并 | 按顺序合并多个 PDF（pdfcpu） |
| PDF 分割 | 每页 / 页码范围 / 每 N 页（pdfcpu） |
| PDF 压缩 | 结构优化减小体积（pdfcpu optimize） |
| PDF 转图片 | 提取 PDF 内嵌图片并导出 JPG / PNG（pdfcpu ExtractImages） |
| 图片转 PDF | JPG / PNG / WEBP → 多页 PDF（pdfcpu ImportImages） |

## 开发

```bash
cd ipdf
wails dev
```

构建：

```bash
cd ipdf
wails build
```

## 技术栈

- Wails v2 + Vue 3 + TypeScript + vue-router
- [pdfcpu](https://github.com/pdfcpu/pdfcpu)：全部 PDF 处理（纯 Go，无 CGO）

## 限制（第一期）

- 不支持加密 PDF 的解密与处理
- 压缩仅为对象级优化，效果中等（后续可接 Ghostscript）
- **PDF 转图片**受 pdfcpu 能力限制：只能提取嵌入图片，**不能**把文字页整页栅格化（扫描件 / 图片合成 PDF 适用）
- 不含 Office / OCR / 水印 / 签署等（首页「即将推出」占位）
- 所有处理均在本地完成，文件不会上传

## 目录

```
ipdf/
  app.go                 # Wails 绑定
  internal/pdf/          # 工具实现
  internal/job/          # 临时工作区
  frontend/src/views/    # 工具 UI
```
