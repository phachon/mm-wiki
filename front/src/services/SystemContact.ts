import httpRequest from './http'
import { ContactEditResp, ContactListResp } from '../types/contactType'
import Base from './Base'

const systemContactUrl = {
  contactSave: '/system/contact/save',
  contactEdit: '/system/contact/edit',
  contactModify: '/system/contact/modify',
  contactList: '/system/contact/list',
  contactDelete: '/system/contact/delete'
}

class SystemContact extends Base {
  public constructor() {
    super()
  }

  public saveContact(contactInfo: {}): Promise<any> {
    const saveContactUrl = this.getProxyUrl(systemContactUrl.contactSave)
    return httpRequest.post<any>(saveContactUrl, {}, contactInfo)
  }

  public getEditContactInfo(contactId: number): Promise<ContactEditResp> {
    const contactEditUrl = this.getProxyUrl(systemContactUrl.contactEdit)
    return httpRequest.get<ContactEditResp>(contactEditUrl, {
      contact_id: contactId
    })
  }

  public modifyContact(editContactInfo: {}): Promise<any> {
    const contactModifyUrl = this.getProxyUrl(systemContactUrl.contactModify)
    return httpRequest.post<any>(contactModifyUrl, {}, editContactInfo)
  }

  public getContactList(
    pageSize?: number,
    pageNum?: number,
    keywords?: {}
  ): Promise<ContactListResp> {
    const contactListUrl = this.getProxyUrl(systemContactUrl.contactList)
    return httpRequest.get<ContactListResp>(contactListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  public deleteContact(contactId: number): Promise<any> {
    const deleteContactUrl = this.getProxyUrl(systemContactUrl.contactDelete)
    return httpRequest.post<any>(deleteContactUrl, {}, { contact_id: contactId })
  }
}

export const SystemContactService = new SystemContact()
