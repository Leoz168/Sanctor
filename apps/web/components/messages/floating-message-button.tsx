"use client";

import { useEffect, useRef, useState } from "react";
import type { PointerEvent as ReactPointerEvent } from "react";
import { Grip, Maximize2, MessageCircle, Minimize2, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { ChatWindow } from "@/components/messages/chat-window";
import { ContactsList } from "@/components/messages/contacts-list";
import {
  clearStoredAuthToken,
  getCurrentUser,
  getStoredAuthToken,
  type CurrentUser,
} from "@/lib/auth-client";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export type Contact = {
  id: string;
  name: string;
  avatar: string;
  online: boolean;
  peerUserId?: string;
  lastMessage?: string;
  lastMessageTime?: string;
};

export type Message = {
  id: string;
  senderId: string;
  text: string;
  timestamp: string;
};

type ConversationSummary = {
  groupId: string;
  peerUserId: string;
  lastMessage?: string;
  lastMessageTime?: string;
};

type DirectMessageResponse = {
  id: string;
  groupId: string;
  userId: string;
  content: string;
  messageTime: string;
};

const defaultPanelSize = { width: 380, height: 600 };
const expandedPanelSize = { width: 760, height: 720 };
const minPanelSize = { width: 320, height: 420 };
const maxPanelSize = { width: 900, height: 760 };
const panelMargin = 24;
const dragSensitivity = 1;
const fallbackAvatar = "/images/community-1.jpg";

type Point = {
  x: number;
  y: number;
};

type PanelSize = {
  width: number;
  height: number;
};

type PanelLayout = Point & PanelSize;
type ResizeDirection = "n" | "s" | "e" | "w" | "ne" | "nw" | "se" | "sw";

export function FloatingMessageButton() {
  const [isOpen, setIsOpen] = useState(false);
  const [isExpanded, setIsExpanded] = useState(false);
  const [currentUser, setCurrentUser] = useState<CurrentUser | null>(null);
  const [selectedContact, setSelectedContact] = useState<Contact | null>(null);
  const [contacts, setContacts] = useState<Contact[]>([]);
  const [messagesByGroup, setMessagesByGroup] = useState<Record<string, Message[]>>({});
  const [contactsError, setContactsError] = useState<string | null>(null);
  const [chatError, setChatError] = useState<string | null>(null);
  const [isLoadingContacts, setIsLoadingContacts] = useState(false);
  const [isLoadingMessages, setIsLoadingMessages] = useState(false);
  const [isSending, setIsSending] = useState(false);
  const [layout, setLayout] = useState<PanelLayout>(() =>
    getDefaultLayout(defaultPanelSize),
  );

  const panelRef = useRef<HTMLDivElement | null>(null);
  const liveLayoutRef = useRef<PanelLayout>(layout);

  useEffect(() => {
    if (!isOpen) {
      return;
    }

    void loadConversations();
  }, [isOpen]);

  const applyLayout = (nextLayout: PanelLayout) => {
    liveLayoutRef.current = nextLayout;

    if (!panelRef.current) {
      return;
    }

    panelRef.current.style.left = `${nextLayout.x}px`;
    panelRef.current.style.top = `${nextLayout.y}px`;
    panelRef.current.style.width = `${nextLayout.width}px`;
    panelRef.current.style.height = `${nextLayout.height}px`;
  };

  const commitLayout = (nextLayout = liveLayoutRef.current) => {
    liveLayoutRef.current = nextLayout;
    setLayout(nextLayout);
    setIsExpanded(
      nextLayout.width > defaultPanelSize.width ||
        nextLayout.height > defaultPanelSize.height,
    );
  };

  const openMessages = () => {
    const nextLayout = getDefaultLayout(defaultPanelSize);

    liveLayoutRef.current = nextLayout;
    setLayout(nextLayout);
    setIsExpanded(false);
    setIsOpen(true);
  };

  const handleClose = () => {
    const nextLayout = getDefaultLayout(defaultPanelSize);

    setIsOpen(false);
    setIsExpanded(false);
    setSelectedContact(null);
    setContactsError(null);
    setChatError(null);
    setLayout(nextLayout);
    liveLayoutRef.current = nextLayout;
  };

  const toggleExpanded = () => {
    const current = liveLayoutRef.current;
    const targetSize = isExpanded ? defaultPanelSize : expandedPanelSize;
    const target = normalizeSize(targetSize);
    const nextLayout = normalizeLayout({
      x: current.x + current.width - target.width,
      y: current.y + current.height - target.height,
      width: target.width,
      height: target.height,
    });

    applyLayout(nextLayout);
    commitLayout(nextLayout);
  };

  const startDrag = (event: ReactPointerEvent<HTMLButtonElement>) => {
    event.preventDefault();

    const startPointer = { x: event.clientX, y: event.clientY };
    const startLayout = liveLayoutRef.current;
    let nextLayout = startLayout;

    if (panelRef.current) {
      panelRef.current.style.willChange = "transform";
    }

    const movePanel = (moveEvent: PointerEvent) => {
      const dx = (moveEvent.clientX - startPointer.x) * dragSensitivity;
      const dy = (moveEvent.clientY - startPointer.y) * dragSensitivity;

      nextLayout = normalizeLayout({
        ...startLayout,
        x: startLayout.x + dx,
        y: startLayout.y + dy,
      });

      liveLayoutRef.current = nextLayout;

      if (!panelRef.current) {
        return;
      }

      panelRef.current.style.transform = `translate3d(${
        nextLayout.x - startLayout.x
      }px, ${nextLayout.y - startLayout.y}px, 0)`;
    };

    const stopDrag = () => {
      if (panelRef.current) {
        panelRef.current.style.transform = "";
        panelRef.current.style.willChange = "";
      }

      applyLayout(nextLayout);
      commitLayout(nextLayout);
      window.removeEventListener("pointermove", movePanel);
      window.removeEventListener("pointerup", stopDrag);
      window.removeEventListener("pointercancel", stopDrag);
    };

    window.addEventListener("pointermove", movePanel);
    window.addEventListener("pointerup", stopDrag);
    window.addEventListener("pointercancel", stopDrag);
  };

  const startResize = (
    direction: ResizeDirection,
    event: ReactPointerEvent<HTMLButtonElement>,
  ) => {
    event.preventDefault();
    event.stopPropagation();

    const startPointer = { x: event.clientX, y: event.clientY };
    const startLayout = liveLayoutRef.current;

    const resizePanel = (moveEvent: PointerEvent) => {
      const dx = moveEvent.clientX - startPointer.x;
      const dy = moveEvent.clientY - startPointer.y;

      applyLayout(resizeLayout(startLayout, direction, dx, dy));
    };

    const stopResize = () => {
      commitLayout();
      window.removeEventListener("pointermove", resizePanel);
      window.removeEventListener("pointerup", stopResize);
      window.removeEventListener("pointercancel", stopResize);
    };

    window.addEventListener("pointermove", resizePanel);
    window.addEventListener("pointerup", stopResize);
    window.addEventListener("pointercancel", stopResize);
  };

  const loadConversations = async () => {
    const token = getStoredAuthToken();
    if (!token) {
      setContactsError("Log in to view your direct messages.");
      return;
    }

    setIsLoadingContacts(true);
    setContactsError(null);

    try {
      const [user, response] = await Promise.all([
        getCurrentUser(),
        fetch(`${apiBase}/api/dm/groups`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }),
      ]);

      if (response.status === 401 || !user) {
        clearStoredAuthToken();
        setContacts([]);
        setCurrentUser(null);
        setContactsError("Your session expired. Please log in again.");
        return;
      }

      const summaries = (await response.json().catch(() => [])) as ConversationSummary[];
      if (!response.ok) {
        throw new Error("Could not load conversations.");
      }

      const contactsWithUsers = await Promise.all(
        summaries.map(async (summary) => {
          let peerName = "Direct conversation";
          let peerAvatar = fallbackAvatar;

          if (summary.peerUserId) {
            const peerResponse = await fetch(
              `${apiBase}/api/users/get?id=${summary.peerUserId}`,
              {
                headers: {
                  Authorization: `Bearer ${token}`,
                },
              },
            );

            if (peerResponse.ok) {
              const peerUser = (await peerResponse.json()) as Partial<CurrentUser>;
              peerName = peerUser.username || peerName;
              peerAvatar = peerUser.avatar || peerAvatar;
            }
          }

          return {
            id: summary.groupId,
            peerUserId: summary.peerUserId,
            name: peerName,
            avatar: peerAvatar,
            online: false,
            lastMessage: summary.lastMessage || "No messages yet",
            lastMessageTime: summary.lastMessageTime
              ? formatConversationTime(summary.lastMessageTime)
              : "",
          } satisfies Contact;
        }),
      );

      setCurrentUser(user);
      setContacts(contactsWithUsers);
    } catch (error) {
      setContactsError(
        error instanceof Error ? error.message : "Could not load conversations.",
      );
    } finally {
      setIsLoadingContacts(false);
    }
  };

  const loadMessages = async (contact: Contact) => {
    const token = getStoredAuthToken();
    if (!token) {
      setChatError("Log in to view messages.");
      return;
    }

    setSelectedContact(contact);
    setIsLoadingMessages(true);
    setChatError(null);

    try {
      const response = await fetch(
        `${apiBase}/api/dm/messages?groupId=${contact.id}&limit=100`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        },
      );

      if (response.status === 401) {
        clearStoredAuthToken();
        setChatError("Your session expired. Please log in again.");
        return;
      }

      const payload = (await response.json().catch(() => [])) as DirectMessageResponse[];
      if (!response.ok) {
        throw new Error("Could not load messages.");
      }

      setMessagesByGroup((current) => ({
        ...current,
        [contact.id]: payload.map(mapMessage),
      }));
    } catch (error) {
      setChatError(error instanceof Error ? error.message : "Could not load messages.");
    } finally {
      setIsLoadingMessages(false);
    }
  };

  const refreshSelectedConversation = async () => {
    if (!selectedContact) {
      return;
    }

    await Promise.all([loadMessages(selectedContact), loadConversations()]);
  };

  const handleSend = async (content: string) => {
    if (!selectedContact) {
      return;
    }

    const token = getStoredAuthToken();
    if (!token) {
      setChatError("Log in to send messages.");
      return;
    }

    setIsSending(true);
    setChatError(null);

    try {
      const response = await fetch(`${apiBase}/api/dm/messages/send`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          groupId: selectedContact.id,
          content,
        }),
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Could not send message.");
      }

      await refreshSelectedConversation();
    } catch (error) {
      setChatError(error instanceof Error ? error.message : "Could not send message.");
      throw error;
    } finally {
      setIsSending(false);
    }
  };

  const handleUpdateMessage = async (messageId: string, content: string) => {
    const token = getStoredAuthToken();
    if (!token) {
      setChatError("Log in to edit messages.");
      return;
    }

    try {
      const response = await fetch(`${apiBase}/api/dm/messages/update?id=${messageId}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ content }),
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Could not update message.");
      }

      await refreshSelectedConversation();
    } catch (error) {
      setChatError(error instanceof Error ? error.message : "Could not update message.");
      throw error;
    }
  };

  const handleDeleteMessage = async (messageId: string) => {
    const token = getStoredAuthToken();
    if (!token) {
      setChatError("Log in to delete messages.");
      return;
    }

    try {
      const response = await fetch(`${apiBase}/api/dm/messages/delete?id=${messageId}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Could not delete message.");
      }

      await refreshSelectedConversation();
    } catch (error) {
      setChatError(error instanceof Error ? error.message : "Could not delete message.");
      throw error;
    }
  };

  const selectedMessages = selectedContact ? messagesByGroup[selectedContact.id] || [] : [];

  return (
    <>
      {!isOpen && (
        <Button
          onClick={openMessages}
          className="fixed bottom-6 right-6 z-50 h-14 w-14 rounded-full bg-brand-orange text-white shadow-lg shadow-brand-orange/30 transition-transform hover:scale-105 hover:bg-orange-600"
          size="icon"
          aria-label="Open messages"
        >
          <MessageCircle className="h-6 w-6" />
        </Button>
      )}

      {isOpen && (
        <div
          ref={panelRef}
          className="fixed z-50 flex flex-col overflow-hidden rounded-2xl border border-border bg-card shadow-2xl"
          style={{
            left: layout.x,
            top: layout.y,
            width: layout.width,
            height: layout.height,
          }}
        >
          <div className="flex items-center justify-between border-b border-border bg-card px-4 py-3">
            <button
              type="button"
              onPointerDown={startDrag}
              className="flex min-w-0 flex-1 cursor-move touch-none select-none items-center gap-2 text-left"
              aria-label="Move messages"
            >
              <Grip className="h-4 w-4 shrink-0 text-muted-foreground" />
              <h2 className="font-semibold text-foreground">Messages</h2>
            </button>

            <div className="flex items-center gap-1">
              <Button
                variant="ghost"
                size="icon"
                onClick={toggleExpanded}
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
                currentUserId={currentUser?.id || ""}
                messages={selectedMessages}
                isLoading={isLoadingMessages}
                isSending={isSending}
                errorMessage={chatError}
                onBack={() => {
                  setSelectedContact(null);
                  setChatError(null);
                }}
                onSend={handleSend}
                onUpdateMessage={handleUpdateMessage}
                onDeleteMessage={handleDeleteMessage}
              />
            ) : (
              <div className="flex h-full flex-col">
                {contactsError ? (
                  <div className="border-b border-border bg-red-50 px-4 py-3 text-sm font-medium text-red-700">
                    {contactsError}
                  </div>
                ) : null}
                {isLoadingContacts ? (
                  <div className="px-4 py-4 text-sm font-medium text-muted-foreground">
                    Loading conversations...
                  </div>
                ) : (
                  <ContactsList contacts={contacts} onSelectContact={(contact) => void loadMessages(contact)} />
                )}
              </div>
            )}
          </div>

          <ResizeHandle direction="n" onResizeStart={startResize} />
          <ResizeHandle direction="s" onResizeStart={startResize} />
          <ResizeHandle direction="e" onResizeStart={startResize} />
          <ResizeHandle direction="w" onResizeStart={startResize} />
          <ResizeHandle direction="ne" onResizeStart={startResize} />
          <ResizeHandle direction="nw" onResizeStart={startResize} />
          <ResizeHandle direction="se" onResizeStart={startResize} />
          <ResizeHandle direction="sw" onResizeStart={startResize} />
        </div>
      )}
    </>
  );
}

interface ResizeHandleProps {
  direction: ResizeDirection;
  onResizeStart: (
    direction: ResizeDirection,
    event: ReactPointerEvent<HTMLButtonElement>,
  ) => void;
}

function ResizeHandle({ direction, onResizeStart }: ResizeHandleProps) {
  const className = getResizeHandleClassName(direction);

  return (
    <button
      type="button"
      className={className}
      onPointerDown={(event) => onResizeStart(direction, event)}
      aria-label={`Resize messages ${direction}`}
    />
  );
}

function mapMessage(message: DirectMessageResponse): Message {
  return {
    id: message.id,
    senderId: message.userId,
    text: message.content,
    timestamp: formatMessageTime(message.messageTime),
  };
}

function formatConversationTime(value: string) {
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return "";
  }

  return new Intl.DateTimeFormat("en-US", {
    month: "short",
    day: "numeric",
    hour: "numeric",
    minute: "2-digit",
  }).format(date);
}

function formatMessageTime(value: string) {
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return "";
  }

  return new Intl.DateTimeFormat("en-US", {
    hour: "numeric",
    minute: "2-digit",
  }).format(date);
}

function getDefaultLayout(size: PanelSize): PanelLayout {
  if (typeof window === "undefined") {
    return { x: panelMargin, y: panelMargin, ...size };
  }

  const normalizedSize = normalizeSize(size);

  return normalizeLayout({
    x: window.innerWidth - normalizedSize.width - panelMargin,
    y: window.innerHeight - normalizedSize.height - panelMargin,
    ...normalizedSize,
  });
}

function normalizeSize(size: PanelSize): PanelSize {
  if (typeof window === "undefined") {
    return size;
  }

  return {
    width: Math.min(
      Math.max(minPanelSize.width, size.width),
      Math.min(maxPanelSize.width, window.innerWidth - panelMargin * 2),
    ),
    height: Math.min(
      Math.max(minPanelSize.height, size.height),
      Math.min(maxPanelSize.height, window.innerHeight - panelMargin * 2),
    ),
  };
}

function normalizeLayout(layout: PanelLayout): PanelLayout {
  const size = normalizeSize(layout);
  const maxX = Math.max(panelMargin, window.innerWidth - size.width - panelMargin);
  const maxY = Math.max(panelMargin, window.innerHeight - size.height - panelMargin);

  return {
    x: Math.min(Math.max(panelMargin, layout.x), maxX),
    y: Math.min(Math.max(panelMargin, layout.y), maxY),
    ...size,
  };
}

function resizeLayout(
  startLayout: PanelLayout,
  direction: ResizeDirection,
  dx: number,
  dy: number,
): PanelLayout {
  const nextLayout = { ...startLayout };

  if (direction.includes("e")) {
    nextLayout.width = startLayout.width + dx;
  }

  if (direction.includes("s")) {
    nextLayout.height = startLayout.height + dy;
  }

  if (direction.includes("w")) {
    nextLayout.width = startLayout.width - dx;
    nextLayout.x = startLayout.x + dx;
  }

  if (direction.includes("n")) {
    nextLayout.height = startLayout.height - dy;
    nextLayout.y = startLayout.y + dy;
  }

  const normalizedSize = normalizeSize(nextLayout);

  if (direction.includes("w")) {
    nextLayout.x = startLayout.x + startLayout.width - normalizedSize.width;
  }

  if (direction.includes("n")) {
    nextLayout.y = startLayout.y + startLayout.height - normalizedSize.height;
  }

  return normalizeLayout({
    ...nextLayout,
    ...normalizedSize,
  });
}

function getResizeHandleClassName(direction: ResizeDirection) {
  const base =
    "absolute z-10 touch-none rounded-full bg-transparent transition-colors hover:bg-brand-orange/20";

  const classes: Record<ResizeDirection, string> = {
    n: "left-5 right-5 top-0 h-2 cursor-n-resize",
    s: "bottom-0 left-5 right-5 h-2 cursor-s-resize",
    e: "bottom-5 right-0 top-5 w-2 cursor-e-resize",
    w: "bottom-5 left-0 top-5 w-2 cursor-w-resize",
    ne: "right-0 top-0 h-5 w-5 cursor-ne-resize",
    nw: "left-0 top-0 h-5 w-5 cursor-nw-resize",
    se: "bottom-0 right-0 h-6 w-6 cursor-se-resize",
    sw: "bottom-0 left-0 h-5 w-5 cursor-sw-resize",
  };

  return `${base} ${classes[direction]}`;
}
