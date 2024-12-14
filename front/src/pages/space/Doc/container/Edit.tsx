import { useGlobalStore } from '@/stores'
import DocEditUI from '../component/EditUI'
import { SpaceDocService } from '@/services/SpaceDoc'
import { DocContentSaveReq, DocContentSaveResp } from '@/types/docType'
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
        window.location.href = `/doc/${values.doc_id}`
      })
    })
  }

  return (
    <div>
      <DocEditUI content={store.content} docInfo={store.viewDocInfo} onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default DocEdit
