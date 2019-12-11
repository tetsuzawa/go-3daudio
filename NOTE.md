
# MEMO

- HRTFを角度ごとにAPIをつくる
- three.jsで音を流す
- トーナメント式で個人にフィット
  - DNN
  - 論文
- APIはgRPC
- API設計より先にAPIの基礎の作成
- データべ0すはMongodb
- サービス間はgRPC
- APIはjson
  - procol buffers
- dataをどのようにjsonで渡すか
  - ストレージのパスを渡す
  - ダウンロードにするか
  - 直接使用するのか
  - フロントエンドとの兼ね合い
  - cdn
  - 速度はどうなのか

- dataは消すのか
- sofaを渡すのか
- wavを渡すのか
- binを渡すのか
- 

- ファイル配信
  - パス
  - キー uuid
  - 複数ファイルならディレクトリ?
  - ディレクトリならzip?

- go でファイルサーバー 

- microservice 案
  - pythonでsofa可視化
  - pythonでグラフ可視化
  - pythonで信号処理
  - golangでhrtf再生
  - three.jsで描画
  - golangでhrtfのフィッティン
  - fileuploader は fork ([uploader](https://github.com/Code-Hex/uploader/))
    - uploaderとpython可視化は同居？
