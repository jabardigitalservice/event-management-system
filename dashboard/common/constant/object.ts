interface IObject {
    key: string,
    label?: string,
    sortable?:boolean
}
export const objectHeaders: IObject[] = [
    {label: "Nama Objek", key: "name", sortable:true},
    {label: "Pengelola", key: "organizer", sortable:true},
    {label: "Alamat Objek", key: "address"},
    {label: "Status", key: "status"},
    {key: "actions"}
]
