interface IObject {
    key: string,
    label?: string,
    sortable?:boolean
}
export const objectHeaders: IObject[] = [
    {label: "Nama Objek Wisata", key: "name", sortable:true},
    {label: "Pengelola", key: "organizer", sortable:true},
    {label: "Status", key: "status"},
    {label: "Aksi",key: "actions"}
]
