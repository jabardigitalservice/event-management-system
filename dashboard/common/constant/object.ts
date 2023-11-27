interface IObject {
    key: string,
    label?: string,
    sortable?:boolean
}
export const objectHeaders: IObject[] = [
    { label: "No", key: "no" },
    {label: "Nama Objek Wisata", key: "name"},
    {label: "Pengelola", key: "organizer"},
    {label: "Status", key: "status"},
    {label: "Aksi",key: "actions"}
]
