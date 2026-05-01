"use client";

import Image from "next/image";
import { Search } from "lucide-react";
import { Input } from "@/components/ui/input";
import type { Contact } from "@/components/messages/floating-message-button";

interface ContactsListProps {
  contacts: Contact[];
  onSelectContact: (contact: Contact) => void;
}

export function ContactsList({ contacts, onSelectContact }: ContactsListProps) {
  return (
    <div className="flex h-full flex-col">
      <div className="px-4 py-3">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Search conversations..."
            className="border-0 bg-muted/50 pl-9"
          />
        </div>
      </div>

      <div className="flex-1 overflow-y-auto">
        {contacts.map((contact) => (
          <button
            key={contact.id}
            onClick={() => onSelectContact(contact)}
            className="flex w-full items-center gap-3 px-4 py-3 text-left transition-colors hover:bg-muted/50"
          >
            <div className="relative flex-shrink-0">
              <div className="h-12 w-12 overflow-hidden rounded-full bg-muted">
                <Image
                  src={contact.avatar}
                  alt={contact.name}
                  width={48}
                  height={48}
                  className="h-full w-full object-cover"
                />
              </div>
              {contact.online && (
                <span className="absolute bottom-0 right-0 h-3 w-3 rounded-full border-2 border-card bg-green-500" />
              )}
            </div>

            <div className="min-w-0 flex-1">
              <div className="flex items-center justify-between">
                <span className="truncate font-medium text-foreground">
                  {contact.name}
                </span>
                <span className="ml-2 flex-shrink-0 text-xs text-muted-foreground">
                  {contact.lastMessageTime}
                </span>
              </div>
              <p className="truncate text-sm text-muted-foreground">
                {contact.lastMessage}
              </p>
            </div>
          </button>
        ))}
      </div>
    </div>
  );
}
