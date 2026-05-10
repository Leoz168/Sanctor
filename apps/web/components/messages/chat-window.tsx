"use client";

import { useState } from "react";
import Image from "next/image";
import { ArrowLeft, MoreVertical, Pencil, Plus, Send, Smile, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import type { Contact, Message } from "@/components/messages/floating-message-button";

interface ChatWindowProps {
  contact: Contact;
  currentUserId: string;
  messages: Message[];
  isLoading?: boolean;
  isSending?: boolean;
  errorMessage?: string | null;
  onBack: () => void;
  onSend: (content: string) => Promise<void>;
  onUpdateMessage: (messageId: string, content: string) => Promise<void>;
  onDeleteMessage: (messageId: string) => Promise<void>;
}

export function ChatWindow({
  contact,
  currentUserId,
  messages,
  isLoading = false,
  isSending = false,
  errorMessage,
  onBack,
  onSend,
  onUpdateMessage,
  onDeleteMessage,
}: ChatWindowProps) {
  const [inputValue, setInputValue] = useState("");

  const handleSend = async () => {
    const trimmed = inputValue.trim();
    if (!trimmed) {
      return;
    }

    await onSend(trimmed);
    setInputValue("");
  };

  const handleEdit = async (message: Message) => {
    const nextContent = window.prompt("Edit message", message.text);
    if (nextContent == null) {
      return;
    }

    const trimmed = nextContent.trim();
    if (!trimmed || trimmed === message.text) {
      return;
    }

    await onUpdateMessage(message.id, trimmed);
  };

  const handleDelete = async (messageId: string) => {
    if (!window.confirm("Delete this message?")) {
      return;
    }

    await onDeleteMessage(messageId);
  };

  return (
    <div className="flex h-full flex-col">
      <div className="flex items-center gap-3 border-b border-border px-3 py-2">
        <Button
          variant="ghost"
          size="icon"
          onClick={onBack}
          className="h-8 w-8 flex-shrink-0"
        >
          <ArrowLeft className="h-4 w-4" />
        </Button>

        <div className="relative flex-shrink-0">
          <div className="h-10 w-10 overflow-hidden rounded-full bg-muted">
            <Image
              src={contact.avatar}
              alt={contact.name}
              width={40}
              height={40}
              className="h-full w-full object-cover"
            />
          </div>
          {contact.online && (
            <span className="absolute bottom-0 right-0 h-2.5 w-2.5 rounded-full border-2 border-card bg-green-500" />
          )}
        </div>

        <div className="min-w-0 flex-1">
          <h3 className="text-sm font-semibold text-foreground">{contact.name}</h3>
          <span className="text-xs text-muted-foreground">
            {contact.lastMessageTime || "Conversation"}
          </span>
        </div>

        <Button variant="ghost" size="icon" className="h-8 w-8 flex-shrink-0">
          <MoreVertical className="h-4 w-4" />
        </Button>
      </div>

      <div className="flex-1 overflow-y-auto px-4 py-4">
        {isLoading ? (
          <div className="rounded-2xl bg-muted/40 px-4 py-6 text-sm font-medium text-muted-foreground">
            Loading messages...
          </div>
        ) : messages.length === 0 ? (
          <div className="rounded-2xl bg-muted/40 px-4 py-6 text-sm font-medium text-muted-foreground">
            No messages yet. Start the conversation.
          </div>
        ) : (
          <div className="space-y-4">
            <div className="flex justify-center">
              <span className="rounded-full bg-muted px-3 py-1 text-xs text-muted-foreground">
                Messages
              </span>
            </div>

            {messages.map((message) => {
              const isMe = message.senderId === currentUserId;

              return (
                <div key={message.id} className={`flex ${isMe ? "justify-end" : "justify-start"}`}>
                  <div className="max-w-[85%]">
                    <div className="flex items-start gap-2">
                      {!isMe && (
                        <div className="h-8 w-8 flex-shrink-0 overflow-hidden rounded-full bg-muted">
                          <Image
                            src={contact.avatar}
                            alt={contact.name}
                            width={32}
                            height={32}
                            className="h-full w-full object-cover"
                          />
                        </div>
                      )}
                      <div>
                        <div
                          className={`rounded-2xl px-4 py-3 ${
                            isMe
                              ? "rounded-br-md bg-primary text-primary-foreground"
                              : "rounded-bl-md bg-muted text-foreground"
                          }`}
                        >
                          <p className="text-sm leading-relaxed">{message.text}</p>
                        </div>

                        {isMe ? (
                          <div className="mt-2 flex justify-end gap-2">
                            <button
                              type="button"
                              onClick={() => handleEdit(message)}
                              className="inline-flex items-center gap-1 text-[11px] font-semibold text-muted-foreground transition-colors hover:text-foreground"
                            >
                              <Pencil className="h-3 w-3" />
                              Edit
                            </button>
                            <button
                              type="button"
                              onClick={() => handleDelete(message.id)}
                              className="inline-flex items-center gap-1 text-[11px] font-semibold text-muted-foreground transition-colors hover:text-red-600"
                            >
                              <Trash2 className="h-3 w-3" />
                              Delete
                            </button>
                          </div>
                        ) : null}
                      </div>
                    </div>
                    <p className={`mt-1 text-xs text-muted-foreground ${isMe ? "text-right" : "ml-10"}`}>
                      {message.timestamp}
                    </p>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>

      {errorMessage ? (
        <div className="border-t border-border bg-red-50 px-3 py-2 text-sm font-medium text-red-700">
          {errorMessage}
        </div>
      ) : null}

      <div className="border-t border-border bg-card px-3 py-3">
        <div className="flex items-center gap-2">
          <Button variant="ghost" size="icon" className="h-9 w-9 flex-shrink-0 text-muted-foreground">
            <Plus className="h-5 w-5" />
          </Button>

          <Input
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            placeholder="Type a message..."
            className="flex-1 border-0 bg-muted/50"
            onKeyDown={(e) => {
              if (e.key === "Enter" && !isSending) {
                void handleSend();
              }
            }}
          />

          <Button variant="ghost" size="icon" className="h-9 w-9 flex-shrink-0 text-muted-foreground">
            <Smile className="h-5 w-5" />
          </Button>

          <Button
            size="icon"
            className="h-9 w-9 rounded-full bg-primary hover:bg-primary/90"
            onClick={() => void handleSend()}
            disabled={isSending}
          >
            <Send className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
