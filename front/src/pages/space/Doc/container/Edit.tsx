import { useGlobalStore } from '@/stores'
import DocEditUI from '../component/EditUI'
import { SpaceDocService } from '@/services/SpaceDoc'
import { DocContentSaveReq } from '@/types/docType'
import { message } from 'antd'
import { useNavigate } from 'react-router-dom'

const DocEdit: React.FC = () => {
  const store = useGlobalStore()
  const navigate = useNavigate()

  const onSaveSubmit = async (values: any) => {
    const req: DocContentSaveReq = {
      doc_id: values.doc_id,
      content: values.content,
      name: values.name
    }
    SpaceDocService.saveDocContent(req).then(() => {
      message.success('保存成功', 2).then(() => {
        store.initDocsByDocId(values.doc_id)
        navigate(`/doc/${values.doc_id}`)
      })
    })
  }

  // onFileUpload 文件上传
  const onFileUpload = (file: any, callback: any, docId?: number) => {
    if (!docId) {
      message.error('上传参数异常')
      return
    }
    SpaceDocService.uploadFile(file, {
      doc_id: docId
    })
      .then((resp) => {
        callback(resp.url) // 上传成功后回调返回文件地址 url
      })
      .catch((e) => {
        console.error('上传失败', e)
      })
  }

  return (
    <div>
      <DocEditUI
        content={store.content}
        docInfo={store.viewDocInfo}
        onSaveSubmit={onSaveSubmit}
        onFileUpload={onFileUpload}
      />
    </div>
  )
}

export default DocEdit
