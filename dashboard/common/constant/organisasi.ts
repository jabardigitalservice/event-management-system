interface IOrganizationHeader {
    key: string,
    label?: string,
    sortable?: boolean
  }
  
export const organisasiHeaders: IOrganizationHeader[] = [
    { label: "Organisasi", key: "name", sortable: true },
    { label: "Email", key: "email", sortable: true },
    { label: "Address", key: "address" },
    { label: "Photo", key: "photo" },
    { key: 'actions' }
];