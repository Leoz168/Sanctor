"use client";

import { useState } from "react";
import Image from "next/image";
import { ArrowLeft, ExternalLink, MoreVertical, Plus, Send, Smile } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import type { Contact, Message } from "@/components/messages/floating-message-button";

interface ChatWindowProps {
  contact: Contact;
  messages: Message[];
  onBack: () => void;
}

export function ChatWindow({ contact, messages, onBack }: ChatWindowProps) {
  const [inputValue, setInputValue] = useState("");

  const handleSend = () => {
    if (inputValue.trim()) {
      setInputValue("");
    }
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
          {contact.online && (
            <span className="text-xs font-medium text-green-500">ONLINE</span>
          )}
        </div>

        <Button variant="ghost" size="icon" className="h-8 w-8 flex-shrink-0">
          <MoreVertical className="h-4 w-4" />
        </Button>
      </div>

      <div className="flex-1 space-y-4 overflow-y-auto px-4 py-4">
        <div className="flex justify-center">
          <span className="rounded-full bg-muted px-3 py-1 text-xs text-muted-foreground">
            Today
          </span>
        </div>

        {messages.map((message) => {
          const isMe = message.senderId === "me";

          if (message.type === "property" && message.property) {
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
                    <div className="overflow-hidden rounded-xl border border-border bg-card p-2">
                      <div className="flex items-center gap-3">
                        <div className="h-16 w-16 flex-shrink-0 overflow-hidden rounded-lg bg-muted">
                          <Image
                            src={message.property.image}
                            alt={message.property.title}
                            width={64}
                            height={64}
                            className="h-full w-full object-cover"
                          />
                        </div>
                        <div className="min-w-0 flex-1">
                          <h4 className="truncate text-sm font-medium text-foreground">
                            {message.property.title}
                          </h4>
                          <p className="text-sm text-muted-foreground">
                            {message.property.price}
                          </p>
                        </div>
                        <Button variant="ghost" size="icon" className="h-8 w-8 flex-shrink-0">
                          <ExternalLink className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            );
          }

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
                  <div
                    className={`rounded-2xl px-4 py-3 ${
                      isMe
                        ? "rounded-br-md bg-primary text-primary-foreground"
                        : "rounded-bl-md bg-muted text-foreground"
                    }`}
                  >
                    <p className="text-sm leading-relaxed">{message.text}</p>
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
            onKeyDown={(e) => e.key === "Enter" && handleSend()}
          />

          <Button variant="ghost" size="icon" className="h-9 w-9 flex-shrink-0 text-muted-foreground">
            <Smile className="h-5 w-5" />
          </Button>

          <Button
            size="icon"
            className="h-9 w-9 rounded-full bg-primary hover:bg-primary/90"
            onClick={handleSend}
          >
            <Send className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
