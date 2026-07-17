import { BellRing, MessageCircleMore, MoreVertical, Paperclip, Phone, Search, Send, Smile, Sparkles, Video } from "lucide-react";

import { ScrollFade } from "@/components/scroll-fade";
import { Button } from "@/components/ui/button";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Bubble, BubbleContent } from "@/components/ui/bubble";
import { Message, MessageAvatar, MessageContent, MessageGroup } from "@/components/ui/message"

const messages = [
    {
        id: 1,
        author: "Alex",
        text: "Привет! Я уже подготовил черновик для нового экрана.",
        time: "09:41",
        own: false,
    },
    {
        id: 2,
        author: "You",
        text: "Отлично, давай посмотрим на него после обеда.",
        time: "09:43",
        own: true,
    },
    {
        id: 3,
        author: "Alex",
        text: "Сделаю ещё пару анимаций, чтобы было ощущение живого мессенджера.",
        time: "09:45",
        own: false,
    },
    {
        id: 4,
        author: "Alex",
        text: "Сделаю ещё пару анимаций, чтобы было ощущение живого мессенджера.",
        time: "09:45",
        own: false,
    },
];

export default function Page() {
    return (
        <div className="min-h-screen bg-[radial-gradient(circle_at_top_left,_rgba(255,255,255,0.04),_transparent_30%),linear-gradient(135deg,_#050505,_#0c0c0f)] text-zinc-100">
            <div className="flex h-screen w-screen overflow-hidden border border-white/10 bg-zinc-950/95 shadow-[0_30px_80px_-30px_rgba(0,0,0,0.8)] backdrop-blur-xl">
                <aside className="hidden w-88 shrink-0 border-r border-white/10 bg-zinc-950/80 p-4 md:flex md:flex-col">
                    <div className="mb-4 flex items-center justify-between rounded-2xl border border-white/10 bg-zinc-900/80 p-3 shadow-lg shadow-black/20">
                        <div>
                            <p className="text-sm font-semibold text-zinc-100">Messenger</p>
                            <p className="text-xs text-zinc-500">Team workspace</p>
                        </div>
                        <Button variant="ghost" size="icon-sm" className="rounded-full text-zinc-300 hover:bg-zinc-800 hover:text-white">
                            <BellRing className="size-4" />
                        </Button>
                    </div>

                    <label className="mb-4 flex items-center gap-2 rounded-2xl border border-white/10 bg-zinc-900/80 px-3 py-2 text-sm text-zinc-400 shadow-inner shadow-black/20">
                        <Search className="size-4" />
                        <input
                            className="w-full bg-transparent outline-none placeholder:text-zinc-500"
                            placeholder="Search chats"
                        />
                    </label>

                    <div className="mb-3 flex items-center justify-between px-1">
                        <p className="text-sm font-semibold text-zinc-200">Recent chats</p>
                        <Button variant="ghost" size="sm" className="rounded-full text-xs text-zinc-400 hover:bg-zinc-800 hover:text-white">
                            <Sparkles className="mr-1 size-3.5" />
                            New
                        </Button>
                    </div>

                    <div className="min-h-0 flex-1 overflow-hidden">
                        <ScrollFade />
                    </div>
                </aside>

                <main className="flex min-w-0 flex-1 flex-col bg-zinc-950">
                    <header className="flex items-center justify-between border-b border-white/10 bg-zinc-950/90 px-4 py-3 sm:px-5">
                        <div className="flex items-center gap-3">
                            <div className="flex size-11 items-center justify-center rounded-2xl bg-gradient-to-br from-violet-600 to-cyan-600 font-semibold text-white shadow-lg shadow-cyan-950/30">
                                A
                            </div>
                            <div>
                                <h2 className="font-semibold text-zinc-100">Alex</h2>
                                <p className="text-sm text-zinc-500">Online • Last seen 2 min ago</p>
                            </div>
                        </div>

                        <div className="flex items-center gap-2">
                            <Button variant="ghost" size="icon-sm" className="rounded-full text-zinc-400 hover:bg-zinc-800 hover:text-white">
                                <Phone className="size-4" />
                            </Button>
                            <Button variant="ghost" size="icon-sm" className="rounded-full text-zinc-400 hover:bg-zinc-800 hover:text-white">
                                <Video className="size-4" />
                            </Button>
                            <Button variant="ghost" size="icon-sm" className="rounded-full text-zinc-400 hover:bg-zinc-800 hover:text-white">
                                <MoreVertical className="size-4" />
                            </Button>
                        </div>
                    </header>

                    <div className="flex-1 overflow-y-auto bg-[linear-gradient(180deg,_rgba(255,255,255,0.03),_rgba(255,255,255,0.01))] p-4 sm:p-5">
                        <div className="mx-auto flex flex-col gap-3">

                            {/* <MessageGroup key={index}> */}
                            {messages.map((message, index) => (
                                <Message key={index} align={message.own ? "end" : "start"}>
                                    <MessageAvatar>
                                        <Avatar>
                                            <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                                            <AvatarFallback>CN</AvatarFallback>
                                        </Avatar>
                                    </MessageAvatar>
                                    <MessageContent>
                                        <Bubble>
                                            <BubbleContent>How can I help you today?</BubbleContent>
                                        </Bubble>
                                    </MessageContent>
                                </Message>
                            ))}
                            {/* </MessageGroup> */}
                            <div className="flex items-center justify-start">
                                <div className="rounded-2xl border border-dashed border-white/10 bg-zinc-900/70 px-4 py-3 text-sm text-zinc-400">
                                    <div className="flex items-center gap-2">
                                        <MessageCircleMore className="size-4" />
                                        New message ideas are ready to send.
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className="border-t border-white/10 bg-zinc-950/90 p-3 sm:p-4">
                        <div className="mx-auto flex items-center gap-2 rounded-2xl border border-white/10 bg-zinc-900/80 px-2 py-2 shadow-inner shadow-black/20">
                            <Button variant="ghost" size="icon-sm" className="rounded-full text-zinc-400 hover:bg-zinc-800 hover:text-white">
                                <Paperclip className="size-4" />
                            </Button>
                            <input className="min-w-0 flex-1 bg-transparent px-2 py-2 text-sm outline-none placeholder:text-zinc-500" placeholder="Write a message..." />
                            <Button variant="ghost" size="icon-sm" className="rounded-full text-zinc-400 hover:bg-zinc-800 hover:text-white">
                                <Smile className="size-4" />
                            </Button>
                            <Button size="sm" className="rounded-full bg-zinc-100 px-4 text-zinc-950 hover:bg-zinc-200">
                                <Send className="mr-1.5 size-4" />
                            </Button>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
}
