package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MediaUploadWallPaperFile
// media.uploadWallPaperFile owner_id:long file:InputFile mime_type:string admin:Bool = Document;
func (c *MediaCore) MediaUploadWallPaperFile(in *media.TLMediaUploadWallPaperFile) (*mtproto.Document, error) {
	var (
		err      error
		document *mtproto.Document
		file     = in.GetFile()
	)

	if file == nil {
		c.Logger.Errorf("media.uploadWallPaperFile - error: file is nil")
		return nil, mtproto.ErrWallpaperFileInvalid
	}

	document, err = c.svcCtx.Dao.DfsClient.DfsUploadWallPaperFile(c.ctx, &dfs.TLDfsUploadWallPaperFile{
		Creator:  in.OwnerId,
		File:     file,
		MimeType: in.GetMimeType(),
		Admin:    in.GetAdmin(),
	})
	if err != nil {
		c.Logger.Errorf("media.uploadWallPaperFile - error: %v", err)
		err = mtproto.ErrWallpaperFileInvalid
		return nil, err
	}

	if len(document.GetThumbs()) > 0 {
		c.svcCtx.Dao.SavePhotoSizeV2(c.ctx, document.GetId(), document.GetThumbs())
	}
	c.svcCtx.Dao.SaveDocumentV2(c.ctx, file.GetName(), document)

	return document, nil
}
