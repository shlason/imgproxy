# imgproxy

一個可以將儲存在任一地方的圖片進行縮放和套用濾鏡，而不用額外儲存圖片的簡易 proxy server。

在左下方 Original image 輸入欲縮放的 image URL，線上 Demo: [點我](https://imgproxy.sidesideeffect.io/)

![imgproxy demo image](https://raw.githubusercontent.com/shlason/imgproxy/docs/images/demo.png)

## 動機
例如有個網站是專門搜集攝影作品，但攝影作品的圖片檔案普遍都非常大，若每次在前端顯示作品縮圖之類，不需要原圖畫質的情境時，會耗費很多不必要的網路流量且載入慢導致使用體驗相對不佳，這個問題可以在上傳作品時就分成很多不同 size 的圖片來儲存，在不同情境就使用相對應 size，但這樣就會多耗費空間來為每個作品做不同 size 的圖片儲存。

## 解法
imgproxy 可以藉由 QS 來指定特定圖片要取得什麼樣的 size，以及縮放的規則，而不用特別把圖片分為不同 size 一一儲存，不僅不用額外耗費儲存空間也增加了彈性，日後不管要什麼 size 的圖片都可以支援。

因為影像處理的部分，相對都比較耗費機器的效能，所以使用 Cloudflare CDN 來做快取，變成只有在全新的請求時才需處理，用以減緩機器的負擔。

API 使用上，例如我有一個攝影作品 A，在顯示作品 A 的縮圖時可以這樣做：
```
https://imgproxy.sidesideeffect.io/api/image?url={IMAGE_A_URL}&width={W}&height={H}&resize=fit&blur=0

IMAGE_A_URL: 作品 A 原圖的 URL
W:           欲縮放的寬度
H:           欲縮放的高度
並且使用 fit 的方式來 resize 不導致變形及不套用模糊濾鏡
```

## 潛在問題
若同時有多個全新的請求，並且圖片又非常大的情況下，可能機器負擔會非常大，這時候可能可以考慮把這個服務架在 AWS Lambda 上面，並一樣處理 CDN，目前該服務是建在 AWS EC2 上面。

## 技術棧
- Go
- Gin (管理路由和 middleware)
- h2non/bimg (影像處理)
- `gopkg.in/ini.v1` (用來讀取設定檔 configs.ini)
- swaggo/swag (產 API 文件)
- autocert (產 SSL 憑證)
- github Actions (用以 build 專案和部署到 EC2 上)
- Cloudflare DNS
- Cloudflare CDN
- AWS EC2

## 參考
以前做專案時曾遇過類似問題，當初看到 [imgproxy github](https://github.com/imgproxy/imgproxy) 並在當時使用類似的方式解決了，因此想要自己試著做一個簡易的版本看看。
