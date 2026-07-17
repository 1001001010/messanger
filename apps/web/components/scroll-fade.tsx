import { MessageCircle, Phone, Video } from "lucide-react";

const chats = Array.from({ length: 20 }, (_, index) => ({
    id: index + 1,
    name: index % 3 === 0 ? `Nina ${index + 1}` : index % 3 === 1 ? `Oleg ${index + 1}` : `Mila ${index + 1}`,
    status: index % 2 === 0 ? "Online" : "Away",
    preview: index % 2 === 0 ? "Смотрю на макет и правлю детали" : "Отправил новый вариант интерфейса",
    time: `${(index % 12) + 1}m`,
    unread: index % 4 === 0 ? 2 : 0,
}));

export function ScrollFade() {
    return (
        <div className="h-full overflow-y-auto pr-1">
            <div className="space-y-2">
                {chats.map((chat) => (
                    <button
                        key={chat.id}
                        className={`flex w-full items-center gap-3 rounded-2xl border px-3 py-2.5 text-left transition ${chat.unread > 0 ? "border-cyan-500/20 bg-zinc-900/90" : "border-transparent bg-transparent hover:border-white/10 hover:bg-zinc-900/70"}`}
                    >
                        <div className="flex size-10 items-center justify-center rounded-2xl bg-gradient-to-br from-violet-600/80 to-cyan-600/80 text-sm font-semibold text-white">
                            {chat.name[0]}
                        </div>
                        <div className="min-w-0 flex-1">
                            <div className="flex items-center justify-between gap-2">
                                <p className="truncate text-sm font-medium text-zinc-100">{chat.name}</p>
                                <span className="text-xs text-zinc-500">{chat.time}</span>
                            </div>
                            <p className="truncate text-xs text-zinc-400">{chat.preview}</p>
                        </div>
                        <div className="flex flex-col items-end gap-1">
                            {chat.unread > 0 ? (
                                <span className="min-w-5 rounded-full bg-cyan-500/20 px-1.5 py-0.5 text-center text-[10px] font-semibold text-cyan-300">
                                    {chat.unread}
                                </span>
                            ) : (
                                <span className="text-zinc-600">
                                    {chat.id % 2 === 0 ? <Phone className="size-3" /> : <Video className="size-3" />}
                                </span>
                            )}
                            <span className="text-[10px] text-zinc-500">{chat.status}</span>
                        </div>
                    </button>
                ))}
            </div>
            <div className="mt-4 flex items-center gap-2 rounded-2xl border border-dashed border-white/10 bg-zinc-900/60 px-3 py-3 text-sm text-zinc-400">
                <MessageCircle className="size-4" />
                <span>More chats will appear here soon.</span>
            </div>
        </div>
    );
}
