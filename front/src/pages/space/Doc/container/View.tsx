import { useGlobalStore } from '@/stores'
import DocViewUI from '../component/ViewUI'

const DocView: React.FC = () => {
  const store = useGlobalStore()
  return (
    <DocViewUI
      loading={store.viewDocLoading}
      docInfo={store.viewDocInfo}
      content={store.content}
      parentPath={store.parentPath}
    />
  )
}

export default DocView
