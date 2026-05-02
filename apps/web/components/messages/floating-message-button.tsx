"use client";

import { useState } from "react";
import { Maximize2, MessageCircle, Minimize2, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { ContactsList } from "@/components/messages/contacts-list";
import { ChatWindow } from "@/components/messages/chat-window";

export type Contact = {
  id: string;
  name: string;
  avatar: string;
  online: boolean;
  lastMessage?: string;
  lastMessageTime?: string;
};

export type Message = {
  id: string;
  senderId: string;
  text: string;
  timestamp: string;
  type: "text" | "property";
  property?: {
    title: string;
    price: string;
    image: string;
  };
};

const mockContacts: Contact[] = [
  {
    id: "1",
    name: "Alex Rivera",
    avatar: "/images/community-1.jpg",
    online: true,
    lastMessage: "2:00 PM is perfect. Should I bring any...",
    lastMessageTime: "10:27 AM",
  },
  {
    id: "2",
    name: "Sarah Chen",
    avatar: "/images/community-4.jpg",
    online: true,
    lastMessage: "Thanks for the tour yesterday!",
    lastMessageTime: "Yesterday",
  },
  {
    id: "3",
    name: "Marcus Thompson",
    avatar: "/images/community-5.jpg",
    online: false,
    lastMessage: "I'll review the documents and get back...",
    lastMessageTime: "2 days ago",
  },
  {
    id: "4",
    name: "Elena Rodriguez",
    avatar: "/images/listing-4.jpg",
    online: false,
    lastMessage: "The apartment looks great!",
    lastMessageTime: "1 week ago",
  },
];

const mockMessages: Record<string, Message[]> = {
  "1": [
    {
      id: "m1",
      senderId: "1",
      text: "Hey! I just saw the listing for the downtown apartment. Is it still available for a viewing this Saturday?",
      timestamp: "10:24 AM",
      type: "text",
    },
    {
      id: "m2",
      senderId: "me",
      text: "Hi Alex! Yes, it is. We have a slot open at 2:00 PM. Does that work for you?",
      timestamp: "10:26 AM",
      type: "text",
    },
    {
      id: "m3",
      senderId: "1",
      text: "2:00 PM is perfect. Should I bring any specific documents with me?",
      timestamp: "10:27 AM",
      type: "text",
    },
    {
      id: "m4",
      senderId: "1",
      text: "",
      timestamp: "10:28 AM",
      type: "property",
      property: {
        title: "Skyline Loft - Unit 402",
        price: "$2,450 / mo",
        image: "/images/listing-1.jpg",
      },
    },
    {
      id: "m5",
      senderId: "me",
      text: "Just your ID for now. I'll send over the application forms after the viewing.",
      timestamp: "10:30 AM",
      type: "text",
    },
  ],
};

export function FloatingMessageButton() {
  const [isOpen, setIsOpen] = useState(false);
  const [isExpanded, setIsExpanded] = useState(false);
  const [selectedContact, setSelectedContact] = useState<Contact | null>(null);
  const [contacts] = useState<Contact[]>(mockContacts);

  const handleSelectContact = (contact: Contact) => {
    setSelectedContact(contact);
  };

  const handleBack = () => {
    setSelectedContact(null);
  };

  const handleClose = () => {
    setIsOpen(false);
    setIsExpanded(false);
    setSelectedContact(null);
  };

  const chatPanelClassName = isExpanded
    ? "fixed bottom-6 right-6 z-50 flex h-[min(720px,calc(100vh-6rem))] w-[min(760px,calc(100vw-3rem))] flex-col overflow-hidden rounded-2xl border border-border bg-card shadow-2xl animate-in slide-in-from-bottom-4 fade-in duration-200"
    : "fixed bottom-6 right-6 z-50 flex h-[600px] w-[calc(100vw-2rem)] max-w-[380px] flex-col overflow-hidden rounded-2xl border border-border bg-card shadow-2xl animate-in slide-in-from-bottom-4 fade-in duration-200";

  return (
    <>
      <Button
        onClick={() => setIsOpen(true)}
        className="fixed bottom-6 right-6 z-50 h-14 w-14 rounded-full bg-primary shadow-lg transition-transform hover:scale-105 hover:bg-primary/90"
        size="icon"
        aria-label="Open messages"
      >
        <MessageCircle className="h-6 w-6" />
      </Button>

      {isOpen && (
        <div className={chatPanelClassName}>
          <div className="flex items-center justify-between border-b border-border bg-card px-4 py-3">
            <h2 className="font-semibold text-foreground">Messages</h2>
            <div className="flex items-center gap-1">
              <Button
                variant="ghost"
                size="icon"
                onClick={() => setIsExpanded((expanded) => !expanded)}
                className="h-8 w-8 rounded-full"
                aria-label={isExpanded ? "Shrink messages" : "Expand messages"}
              >
                {isExpanded ? (
                  <Minimize2 className="h-4 w-4" />
                ) : (
                  <Maximize2 className="h-4 w-4" />
                )}
              </Button>

              <Button
                variant="ghost"
                size="icon"
                onClick={handleClose}
                className="h-8 w-8 rounded-full"
                aria-label="Close messages"
              >
                <X className="h-4 w-4" />
              </Button>
            </div>
          </div>

          <div className="flex-1 overflow-hidden">
            {selectedContact ? (
              <ChatWindow
                contact={selectedContact}
                messages={mockMessages[selectedContact.id] || []}
                onBack={handleBack}
              />
            ) : (
              <ContactsList
                contacts={contacts}
                onSelectContact={handleSelectContact}
              />
            )}
          </div>
        </div>
      )}
    </>
  );
}
