import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { Event } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import EventComponent from '@/components/event';
import BaseWrapper from '@/wrappers/base';
import Sidebar from '@/components/common/sidebar';
import MainWrapper from '@/wrappers/main';

const Events = () => {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getEvents = () => {
    setLoading(true);
    const URL = `/flags/events?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedEvents = [...events, ...(res.data.events || [])];
          if (addedEvents.length === events.length) setHasMore(false);
          setEvents(addedEvents);
          setPage(prev => prev + 1);
          setLoading(false);
        } else {
          if (res.data.message) Toaster.error(res.data.message, 'error_toaster');
          else {
            Toaster.error(SERVER_ERROR, 'error_toaster');
          }
        }
      })
      .catch(err => {
        Toaster.error(SERVER_ERROR, 'error_toaster');
      });
  };

  useEffect(() => {
    getEvents();
  }, []);

  return (
    <BaseWrapper>
      <Sidebar index={6} />
      <MainWrapper>
    <div className="w-full h-full p-4">
      {loading ? (
        <div className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-full">
          <Loader />
        </div>
      ) : (
        <InfiniteScroll
          className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-screen flex flex-col gap-2 max-lg:px-4 pb-base_padding"
          dataLength={events.length}
          next={getEvents}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {events.length === 0 ? (
            <div>No Flagged events.</div>
          ) : (
            events.map(event => <EventComponent key={event.id} event={event} setEvents={setEvents} />)
          )}
        </InfiniteScroll>
      )}
    </div>
    </MainWrapper>
    </BaseWrapper>
  );
};

export default Events;
