interface IOrganizationHeader {
    key: string,
    label?: string,
    sortable?: boolean
  }
  
export const organisasiHeaders: IOrganizationHeader[] = [
    { label: "No", key: "no" },
    { label: "Organisasi", key: "name" },
    { label: "Email", key: "email" },
    { label: "Address", key: "address" },
    { label: "Aksi", key: 'actions' }
];