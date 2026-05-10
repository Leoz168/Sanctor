"use client";

import Image from "next/image";
import { Search } from "lucide-react";
import { Input } from "@/components/ui/input";
import type {
  Contact,
  SearchUser,
} from "@/components/messages/floating-message-button";

interface ContactsListProps {
  contacts: Contact[];
  searchQuery: string;
  searchResults: SearchUser[];
  isSearching?: boolean;
  onSearchChange: (value: string) => void;
  onSelectContact: (contact: Contact) => void;
  onSelectSearchResult: (user: SearchUser) => void;
}

export function ContactsList({
  contacts,
  searchQuery,
  searchResults,
  isSearching = false,
  onSearchChange,
  onSelectContact,
  onSelectSearchResult,
}: ContactsListProps) {
  return (
    <div className="flex h-full flex-col">
      <div className="px-4 py-3">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            value={searchQuery}
            onChange={(event) => onSearchChange(event.target.value)}
            placeholder="Search users or conversations..."
            className="border-0 bg-muted/50 pl-9"
          />
        </div>
      </div>

      <div className="flex-1 overflow-y-auto">
        {searchQuery.trim() ? (
          <div className="border-b border-border/70 px-4 pb-3">
            <p className="mb-2 text-[11px] font-bold uppercase tracking-[0.14em] text-muted-foreground">
              Search results
            </p>
            {isSearching ? (
              <p className="text-sm text-muted-foreground">Searching users...</p>
            ) : searchResults.length === 0 ? (
              <p className="text-sm text-muted-foreground">No matching users found.</p>
            ) : (
              <div className="space-y-1">
                {searchResults.map((user) => (
                  <button
                    key={user.id}
                    type="button"
                    onClick={() => onSelectSearchResult(user)}
                    className="flex w-full items-center gap-3 rounded-2xl px-3 py-2 text-left transition-colors hover:bg-muted/50"
                  >
                    <div className="h-10 w-10 overflow-hidden rounded-full bg-muted">
                      <Image
                        src={user.avatar}
                        alt={user.name}
                        width={40}
                        height={40}
                        className="h-full w-full object-cover"
                      />
                    </div>
                    <div className="min-w-0 flex-1">
                      <p className="truncate font-medium text-foreground">{user.name}</p>
                      <p className="truncate text-xs text-muted-foreground">
                        Start a direct conversation
                      </p>
                    </div>
                  </button>
                ))}
              </div>
            )}
          </div>
        ) : null}

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
